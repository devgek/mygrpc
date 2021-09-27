package mygrpc

import (
	"context"
	"encoding/json"
	"github.com/devgek/mygrpc/pb/mygrpcpb"
	"net/http"
)

// implements http.DecodeRequestFunc from Go kit
// for HTTP/JSON transport
func decodeLastNameRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req lastNameRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// implements http.EncodeResponseFunc from Go kit
func encodeLastNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// end of HTTP/JSON transport

// for grpc transport

func EncodeGRPCLastNameRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(lastNameRequest)
	return &mygrpcpb.LastNameRequest{First: req.FirstName}, nil
}

func DecodeGRPCLastNameRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*mygrpcpb.LastNameRequest)
	return lastNameRequest{FirstName: req.First}, nil
}

func EncodeGRPCLastNameResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(lastNameResponse)
	return &mygrpcpb.LastNameResponse{Last: res.LastName, Err: res.ErrMsg}, nil
}

func DecodeGRPCLastNameResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*mygrpcpb.LastNameResponse)
	return lastNameResponse{LastName: res.Last, ErrMsg: res.Err}, nil
}

// end of encode/decode grpc transport
