package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/devgek/mygrpc"
	"github.com/devgek/mygrpc/pb/mygrpcpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := mygrpc.NewService()
	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("catched signal written to errChan %s", <-c)
	}()

	getLastNameEndpoint := mygrpc.MakeGetLastNameEndpoint(srv)
	{
		getLastNameEndpoint = mygrpc.NewSimpleLoggingMiddleware()(getLastNameEndpoint)
	}
	endpoints := mygrpc.MyEndpoints{
		GetLastNameEndpoint: getLastNameEndpoint,
	}

	// HTTP transport
	go func() {
		log.Println("http:", *httpAddr)
		handler := mygrpc.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// gRPC transport
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("grpc:", *gRPCAddr)
		handler := mygrpc.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		mygrpcpb.RegisterMyGrpcServer(gRPCServer, handler)

		errChan <- gRPCServer.Serve(listener)
	}()

	log.Fatalln(<-errChan)
}
