package main

import (
	"fmt"
	"log"
	"net"

	"github.com/adarshjeetAmplio/grpc-server/internal/data"
	proto "github.com/adarshjeetAmplio/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedUserServiceServer
}

func main() {
	data.InitDatabase()
	NewGRPC()
}

func NewGRPC() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	fmt.Println("gRPC server running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
