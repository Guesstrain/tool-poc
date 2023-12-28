package main

import (
	"fmt"
	"log"
	"net"

	"tool/grpc-api-poc/chat"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC poc!")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
