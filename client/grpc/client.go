package client

import (
	"github.com/devgek/mygrpc"
	"github.com/devgek/mygrpc/pb/mygrpcpb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// New makes a new mygrpc.Service client.
func New(conn *grpc.ClientConn) mygrpc.Service {
	var lastNameEndpoint = grpctransport.NewClient(
		conn, "MyGrpc", "LastName",
		mygrpc.EncodeGRPCLastNameRequest,
		mygrpc.DecodeGRPCLastNameResponse,
		mygrpcpb.LastNameResponse{},
	).Endpoint()
	return mygrpc.MyEndpoints{
		GetLastNameEndpoint: lastNameEndpoint,
	}
}

func NewClient(conn *grpc.ClientConn) mygrpcpb.MyGrpcClient {
	return mygrpcpb.NewMyGrpcClient(conn)
}
