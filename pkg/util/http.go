package util

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	TraceIDVar                = "traceID"
	AcceptHeaderKey           = "Accept"
	ProtobufTypeHeaderValue   = "application/protobuf"
	JSONTypeHeaderValue       = "application/json"
	JaegerJSONTypeHeaderValue = "application/jaeger-json"
)

func ParseTraceID(r *http.Request) ([]byte, error) {
	vars := mux.Vars(r)
	traceID, ok := vars[TraceIDVar]
	if !ok {
		return nil, fmt.Errorf("please provide a traceID")
	}

	byteID, err := HexStringToTraceID(traceID)
	if err != nil {
		return nil, err
	}

	return byteID, nil
}
