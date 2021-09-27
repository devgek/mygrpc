package mygrpc

import (
	"context"
	"github.com/devgek/mygrpc/pb/mygrpcpb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer ...
func NewGRPCServer(ctx context.Context, endpoints MyEndpoints) mygrpcpb.MyGrpcServer {
	return &grpcServer{
		lastNameHandler: grpctransport.NewServer(
			endpoints.GetLastNameEndpoint,
			DecodeGRPCLastNameRequest,
			EncodeGRPCLastNameResponse,
		),
	}
}

type grpcServer struct {
	lastNameHandler grpctransport.Handler
}

func (gs *grpcServer) LastName(ctx context.Context, request *mygrpcpb.LastNameRequest) (*mygrpcpb.LastNameResponse, error) {
	_, response, err := gs.lastNameHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*mygrpcpb.LastNameResponse), nil
}
