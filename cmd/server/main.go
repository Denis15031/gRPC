package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc-stub/internal/server"

	"google.golang.org/grpc/reflection"
	"grpc-stub/api"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterUserServiceServer(grpcServer, server.NewUserServer())
	api.RegisterChatServiceServer(grpcServer, server.NewChatServer())

	reflection.Register(grpcServer)

	fmt.Println("gRPC server running on :50051")
	log.Fatal(grpcServer.Serve(lis))
}
