package main

import (
	"context"
	pb "grpc-sample-api/helloworld/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HelloworldHandler struct {
	pb.UnimplementedGreeterServer
}

func (h HelloworldHandler) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	port := ":9090"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &HelloworldHandler{})
	reflection.Register(server)

	log.Printf("start Grpc server")
	server.Serve(lis)
}
