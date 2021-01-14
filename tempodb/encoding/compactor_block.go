package encoding

import (
	"bytes"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/grafana/tempo/tempodb/backend"
	"github.com/grafana/tempo/tempodb/encoding/bloom"
)

type CompactorBlock struct {
	compactedMeta *backend.BlockMeta
	inMetas       []*backend.BlockMeta

	bloom *bloom.ShardedBloomFilter

	bufferedObjects int
	appendBuffer    *bytes.Buffer
	appender        Appender
}

func NewCompactorBlock(id uuid.UUID, tenantID string, bloomFP float64, indexDownsample int, metas []*backend.BlockMeta, estimatedObjects int) (*CompactorBlock, error) {
	if len(metas) == 0 {
		return nil, fmt.Errorf("empty block meta list")
	}

	if estimatedObjects <= 0 {
		return nil, fmt.Errorf("must have non-zero positive estimated objects for a reliable bloom filter")
	}

	c := &CompactorBlock{
		compactedMeta: backend.NewBlockMeta(tenantID, id),
		bloom:         bloom.NewWithEstimates(uint(estimatedObjects), bloomFP),
		inMetas:       metas,
	}

	c.appendBuffer = &bytes.Buffer{}
	c.appender = NewBufferedAppender(c.appendBuffer, indexDownsample, estimatedObjects)

	return c, nil
}

func (c *CompactorBlock) AddObject(id ID, object []byte) error {
	err := c.appender.Append(id, object)
	if err != nil {
		return err
	}
	c.bufferedObjects++
	c.compactedMeta.ObjectAdded(id)
	c.bloom.Add(id)
	return nil
}

func (c *CompactorBlock) CurrentBufferLength() int {
	return c.appendBuffer.Len()
}

func (c *CompactorBlock) CurrentBufferedObjects() int {
	return c.bufferedObjects
}

func (c *CompactorBlock) Length() int {
	return c.appender.Length()
}

// FlushBuffer flushes any existing objects to the backend
func (c *CompactorBlock) FlushBuffer(ctx context.Context, tracker backend.AppendTracker, w backend.Writer) (backend.AppendTracker, error) {
	meta := c.BlockMeta()
	tracker, err := w.Append(ctx, nameObjects, meta.BlockID, meta.TenantID, tracker, c.appendBuffer.Bytes())
	if err != nil {
		return nil, err
	}

	c.appendBuffer.Reset()
	c.bufferedObjects = 0

	return tracker, nil
}

// Complete finishes writes the compactor metadata and closes all buffers and appenders
func (c *CompactorBlock) Complete(ctx context.Context, tracker backend.AppendTracker, w backend.Writer) error {
	c.appender.Complete()

	// index
	records := c.appender.Records()
	indexBytes, err := MarshalRecords(records)
	if err != nil {
		return err
	}

	// bloom
	bloomBuffers, err := c.bloom.WriteTo()
	if err != nil {
		return err
	}

	meta := c.BlockMeta()
	err = writeBlock(ctx, w, meta, indexBytes, bloomBuffers)
	if err != nil {
		return err
	}

	return w.CloseAppend(ctx, tracker)
}

func (c *CompactorBlock) BlockMeta() *backend.BlockMeta {
	meta := c.compactedMeta

	meta.StartTime = c.inMetas[0].StartTime
	meta.EndTime = c.inMetas[0].EndTime

	// everything should be correct here except the start/end times which we will get from the passed in metas
	for _, m := range c.inMetas[1:] {
		if m.StartTime.Before(meta.StartTime) {
			meta.StartTime = m.StartTime
		}
		if m.EndTime.After(meta.EndTime) {
			meta.EndTime = m.EndTime
		}
	}

	return meta
}
