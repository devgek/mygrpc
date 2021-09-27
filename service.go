package mygrpc

import (
	"context"
	"errors"
)

type Service interface {
	GetLastName(ctx context.Context, firstName string) (string, error)
}

type mygrpcService struct{}

func NewService() Service {
	return mygrpcService{}
}

func (s mygrpcService) GetLastName(ctx context.Context, firstName string) (string, error) {
	if firstName == "Lionel" {
		return "Messi", nil
	}

	return "", errors.New("unknown person")
}
