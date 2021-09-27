package mygrpc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type lastNameRequest struct {
	FirstName string `json:"firstName"`
}

type lastNameResponse struct {
	LastName string `json:"lastName"`
	ErrMsg   string `json:"err,omitempty"`
}

func MakeGetLastNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(lastNameRequest)
		lastName, err := s.GetLastName(ctx, req.FirstName)
		if err != nil {
			return lastNameResponse{LastName: lastName, ErrMsg: err.Error()}, nil
		}

		return lastNameResponse{LastName: lastName, ErrMsg: ""}, nil
	}
}

type MyEndpoints struct {
	GetLastNameEndpoint endpoint.Endpoint
}

func (e MyEndpoints) GetLastName(ctx context.Context, firstName string) (string, error) {
	req := lastNameRequest{FirstName: firstName}
	resp, err := e.GetLastNameEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	lastNameResponse := resp.(lastNameResponse)
	if lastNameResponse.ErrMsg != "" {
		return "", errors.New(lastNameResponse.ErrMsg)
	}

	return lastNameResponse.LastName, nil
}
