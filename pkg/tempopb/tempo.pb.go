// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tempo.proto

package tempopb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TraceByIDRequest struct {
	TraceID        []byte `protobuf:"bytes,1,opt,name=traceID,proto3" json:"traceID,omitempty"`
	BlockStart     []byte `protobuf:"bytes,2,opt,name=blockStart,proto3" json:"blockStart,omitempty"`
	BlockEnd       []byte `protobuf:"bytes,3,opt,name=blockEnd,proto3" json:"blockEnd,omitempty"`
	QueryIngesters bool   `protobuf:"varint,4,opt,name=queryIngesters,proto3" json:"queryIngesters,omitempty"`
}

func (m *TraceByIDRequest) Reset()         { *m = TraceByIDRequest{} }
func (m *TraceByIDRequest) String() string { return proto.CompactTextString(m) }
func (*TraceByIDRequest) ProtoMessage()    {}
func (*TraceByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b334b194b16825ec, []int{0}
}
func (m *TraceByIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceByIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceByIDRequest.Merge(m, src)
}
func (m *TraceByIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *TraceByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TraceByIDRequest proto.InternalMessageInfo

func (m *TraceByIDRequest) GetTraceID() []byte {
	if m != nil {
		return m.TraceID
	}
	return nil
}

func (m *TraceByIDRequest) GetBlockStart() []byte {
	if m != nil {
		return m.BlockStart
	}
	return nil
}

func (m *TraceByIDRequest) GetBlockEnd() []byte {
	if m != nil {
		return m.BlockEnd
	}
	return nil
}

func (m *TraceByIDRequest) GetQueryIngesters() bool {
	if m != nil {
		return m.QueryIngesters
	}
	return false
}

type TraceByIDResponse struct {
	Trace *Trace `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
}

func (m *TraceByIDResponse) Reset()         { *m = TraceByIDResponse{} }
func (m *TraceByIDResponse) String() string { return proto.CompactTextString(m) }
func (*TraceByIDResponse) ProtoMessage()    {}
func (*TraceByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b334b194b16825ec, []int{1}
}
func (m *TraceByIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceByIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceByIDResponse.Merge(m, src)
}
func (m *TraceByIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *TraceByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TraceByIDResponse proto.InternalMessageInfo

func (m *TraceByIDResponse) GetTrace() *Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

type Trace struct {
	Batches []*v1.ResourceSpans `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches,omitempty"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_b334b194b16825ec, []int{2}
}
func (m *Trace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return m.Size()
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetBatches() []*v1.ResourceSpans {
	if m != nil {
		return m.Batches
	}
	return nil
}

type PushRequest struct {
	Batch *v1.ResourceSpans `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b334b194b16825ec, []int{3}
}
func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return m.Size()
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetBatch() *v1.ResourceSpans {
	if m != nil {
		return m.Batch
	}
	return nil
}

type PushResponse struct {
}

func (m *PushResponse) Reset()         { *m = PushResponse{} }
func (m *PushResponse) String() string { return proto.CompactTextString(m) }
func (*PushResponse) ProtoMessage()    {}
func (*PushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b334b194b16825ec, []int{4}
}
func (m *PushResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushResponse.Merge(m, src)
}
func (m *PushResponse) XXX_Size() int {
	return m.Size()
}
func (m *PushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TraceByIDRequest)(nil), "tempopb.TraceByIDRequest")
	proto.RegisterType((*TraceByIDResponse)(nil), "tempopb.TraceByIDResponse")
	proto.RegisterType((*Trace)(nil), "tempopb.Trace")
	proto.RegisterType((*PushRequest)(nil), "tempopb.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "tempopb.PushResponse")
}

func init() { proto.RegisterFile("tempo.proto", fileDescriptor_b334b194b16825ec) }

var fileDescriptor_b334b194b16825ec = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4d, 0x6b, 0xe2, 0x50,
	0x14, 0xcd, 0x1b, 0x3f, 0x22, 0x37, 0x8e, 0xcc, 0x3c, 0x66, 0x20, 0x93, 0x45, 0x90, 0x30, 0x0c,
	0x81, 0x81, 0x88, 0x19, 0x66, 0xd1, 0x55, 0xa9, 0x68, 0xa9, 0x9b, 0x62, 0x63, 0xff, 0x40, 0x12,
	0x2f, 0x55, 0xaa, 0x49, 0x7c, 0xef, 0x45, 0xf0, 0x5f, 0xf8, 0xb3, 0xba, 0x74, 0xd9, 0x65, 0xd1,
	0x3f, 0x52, 0x7c, 0xcf, 0x48, 0x94, 0x6e, 0xba, 0x7b, 0xe7, 0x9e, 0x7b, 0x6e, 0xce, 0x39, 0x01,
	0x43, 0xe0, 0x22, 0x4b, 0xbd, 0x8c, 0xa5, 0x22, 0xa5, 0xba, 0x04, 0x59, 0x64, 0xb9, 0x69, 0x86,
	0x89, 0xc0, 0x39, 0x2e, 0x50, 0xb0, 0x75, 0x47, 0xb2, 0x1d, 0xc1, 0xc2, 0x18, 0x3b, 0xab, 0xae,
	0x7a, 0x28, 0x89, 0xb3, 0x21, 0xf0, 0xed, 0xf1, 0x80, 0x7b, 0xeb, 0x61, 0x3f, 0xc0, 0x65, 0x8e,
	0x5c, 0x50, 0x13, 0x74, 0xb9, 0x33, 0xec, 0x9b, 0xa4, 0x4d, 0xdc, 0x66, 0x50, 0x40, 0x6a, 0x03,
	0x44, 0xf3, 0x34, 0x7e, 0x1e, 0x8b, 0x90, 0x09, 0xf3, 0x8b, 0x24, 0x4b, 0x13, 0x6a, 0x41, 0x43,
	0xa2, 0x41, 0x32, 0x31, 0x2b, 0x92, 0x3d, 0x61, 0xfa, 0x07, 0x5a, 0xcb, 0x1c, 0xd9, 0x7a, 0x98,
	0x3c, 0x21, 0x17, 0xc8, 0xb8, 0x59, 0x6d, 0x13, 0xb7, 0x11, 0x5c, 0x4c, 0x9d, 0x2b, 0xf8, 0x5e,
	0x72, 0xc4, 0xb3, 0x34, 0xe1, 0x48, 0x7f, 0x43, 0x4d, 0x7a, 0x90, 0x86, 0x0c, 0xbf, 0xe5, 0x1d,
	0xa3, 0x7a, 0x72, 0x35, 0x50, 0xa4, 0x73, 0x0f, 0x35, 0x89, 0xe9, 0x00, 0xf4, 0x28, 0x14, 0xf1,
	0x14, 0xb9, 0x49, 0xda, 0x15, 0xd7, 0xf0, 0xff, 0x7a, 0x67, 0x95, 0xa8, 0xf4, 0x9e, 0x6a, 0x62,
	0xd5, 0xf5, 0x02, 0xe4, 0x69, 0xce, 0x62, 0x1c, 0x67, 0x61, 0xc2, 0x83, 0x42, 0xeb, 0x8c, 0xc0,
	0x18, 0xe5, 0x7c, 0x5a, 0xf4, 0x72, 0x03, 0x35, 0xc9, 0x1c, 0x4d, 0x7c, 0xea, 0xa6, 0x52, 0x3a,
	0x2d, 0x68, 0xaa, 0x8b, 0x2a, 0x97, 0x7f, 0x0d, 0xf5, 0x03, 0x46, 0x46, 0xff, 0x43, 0xf5, 0xf0,
	0xa2, 0x3f, 0x4e, 0xd1, 0x4a, 0x9f, 0xb6, 0x7e, 0x5e, 0x4c, 0x95, 0xdc, 0xd1, 0xfc, 0x31, 0xe8,
	0x0f, 0x39, 0xb2, 0x19, 0x32, 0x7a, 0x07, 0x5f, 0x6f, 0x67, 0xc9, 0xe4, 0x54, 0x1e, 0xfd, 0x75,
	0xde, 0x52, 0xe9, 0x17, 0x5b, 0xd6, 0x47, 0x54, 0x71, 0xb4, 0x67, 0xbe, 0xec, 0x6c, 0xb2, 0xdd,
	0xd9, 0xe4, 0x6d, 0x67, 0x93, 0xcd, 0xde, 0xd6, 0xb6, 0x7b, 0x5b, 0x7b, 0xdd, 0xdb, 0x5a, 0x54,
	0x97, 0x21, 0xff, 0xbd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x69, 0x55, 0x95, 0x67, 0x78, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PusherClient is the client API for Pusher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PusherClient interface {
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
}

type pusherClient struct {
	cc *grpc.ClientConn
}

func NewPusherClient(cc *grpc.ClientConn) PusherClient {
	return &pusherClient{cc}
}

func (c *pusherClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, "/tempopb.Pusher/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PusherServer is the server API for Pusher service.
type PusherServer interface {
	Push(context.Context, *PushRequest) (*PushResponse, error)
}

// UnimplementedPusherServer can be embedded to have forward compatible implementations.
type UnimplementedPusherServer struct {
}

func (*UnimplementedPusherServer) Push(ctx context.Context, req *PushRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}

func RegisterPusherServer(s *grpc.Server, srv PusherServer) {
	s.RegisterService(&_Pusher_serviceDesc, srv)
}

func _Pusher_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tempopb.Pusher/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pusher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tempopb.Pusher",
	HandlerType: (*PusherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Pusher_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tempo.proto",
}

// QuerierClient is the client API for Querier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuerierClient interface {
	FindTraceByID(ctx context.Context, in *TraceByIDRequest, opts ...grpc.CallOption) (*TraceByIDResponse, error)
}

type querierClient struct {
	cc *grpc.ClientConn
}

func NewQuerierClient(cc *grpc.ClientConn) QuerierClient {
	return &querierClient{cc}
}

func (c *querierClient) FindTraceByID(ctx context.Context, in *TraceByIDRequest, opts ...grpc.CallOption) (*TraceByIDResponse, error) {
	out := new(TraceByIDResponse)
	err := c.cc.Invoke(ctx, "/tempopb.Querier/FindTraceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuerierServer is the server API for Querier service.
type QuerierServer interface {
	FindTraceByID(context.Context, *TraceByIDRequest) (*TraceByIDResponse, error)
}

// UnimplementedQuerierServer can be embedded to have forward compatible implementations.
type UnimplementedQuerierServer struct {
}

func (*UnimplementedQuerierServer) FindTraceByID(ctx context.Context, req *TraceByIDRequest) (*TraceByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTraceByID not implemented")
}

func RegisterQuerierServer(s *grpc.Server, srv QuerierServer) {
	s.RegisterService(&_Querier_serviceDesc, srv)
}

func _Querier_FindTraceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServer).FindTraceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tempopb.Querier/FindTraceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServer).FindTraceByID(ctx, req.(*TraceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Querier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tempopb.Querier",
	HandlerType: (*QuerierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindTraceByID",
			Handler:    _Querier_FindTraceByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tempo.proto",
}

func (m *TraceByIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceByIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceByIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QueryIngesters {
		i--
		if m.QueryIngesters {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.BlockEnd) > 0 {
		i -= len(m.BlockEnd)
		copy(dAtA[i:], m.BlockEnd)
		i = encodeVarintTempo(dAtA, i, uint64(len(m.BlockEnd)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BlockStart) > 0 {
		i -= len(m.BlockStart)
		copy(dAtA[i:], m.BlockStart)
		i = encodeVarintTempo(dAtA, i, uint64(len(m.BlockStart)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TraceID) > 0 {
		i -= len(m.TraceID)
		copy(dAtA[i:], m.TraceID)
		i = encodeVarintTempo(dAtA, i, uint64(len(m.TraceID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TraceByIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceByIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceByIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Trace != nil {
		{
			size, err := m.Trace.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTempo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Trace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Trace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Batches) > 0 {
		for iNdEx := len(m.Batches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Batches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTempo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PushRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Batch != nil {
		{
			size, err := m.Batch.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTempo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PushResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTempo(dAtA []byte, offset int, v uint64) int {
	offset -= sovTempo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TraceByIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TraceID)
	if l > 0 {
		n += 1 + l + sovTempo(uint64(l))
	}
	l = len(m.BlockStart)
	if l > 0 {
		n += 1 + l + sovTempo(uint64(l))
	}
	l = len(m.BlockEnd)
	if l > 0 {
		n += 1 + l + sovTempo(uint64(l))
	}
	if m.QueryIngesters {
		n += 2
	}
	return n
}

func (m *TraceByIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Trace != nil {
		l = m.Trace.Size()
		n += 1 + l + sovTempo(uint64(l))
	}
	return n
}

func (m *Trace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Batches) > 0 {
		for _, e := range m.Batches {
			l = e.Size()
			n += 1 + l + sovTempo(uint64(l))
		}
	}
	return n
}

func (m *PushRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Batch != nil {
		l = m.Batch.Size()
		n += 1 + l + sovTempo(uint64(l))
	}
	return n
}

func (m *PushResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTempo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTempo(x uint64) (n int) {
	return sovTempo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TraceByIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TraceByIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceByIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceID = append(m.TraceID[:0], dAtA[iNdEx:postIndex]...)
			if m.TraceID == nil {
				m.TraceID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockStart", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockStart = append(m.BlockStart[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockStart == nil {
				m.BlockStart = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockEnd = append(m.BlockEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockEnd == nil {
				m.BlockEnd = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryIngesters", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.QueryIngesters = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TraceByIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TraceByIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceByIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trace == nil {
				m.Trace = &Trace{}
			}
			if err := m.Trace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Trace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Trace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Batches = append(m.Batches, &v1.ResourceSpans{})
			if err := m.Batches[len(m.Batches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Batch == nil {
				m.Batch = &v1.ResourceSpans{}
			}
			if err := m.Batch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTempo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTempo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTempo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTempo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTempo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTempo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTempo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTempo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTempo = fmt.Errorf("proto: unexpected end of group")
)
