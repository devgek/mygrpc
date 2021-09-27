package main

import (
	"context"
	"flag"
	"fmt"
	client "github.com/devgek/mygrpc/client/grpc"
	"github.com/devgek/mygrpc/pb/mygrpcpb"
	"log"
	"time"

	"google.golang.org/grpc"
)

/*
	Usage

		cli last firtName

*/

func main() {
	var (
		grpcAddr = flag.String("addr", ":8081", "gRPC address")
	)
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()
	//mygrpcService := client.New(conn)
	myClient := client.NewClient(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "last":
		var firstName string
		firstName, args = pop(args)
		getLastName(ctx, myClient, firstName)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

func getLastName(ctx context.Context, client mygrpcpb.MyGrpcClient, firstName string) {
	req := mygrpcpb.LastNameRequest{First: firstName}
	h, err := client.LastName(ctx, &req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(h)
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
