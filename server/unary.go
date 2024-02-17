package main

import (
	"context"
	"log"

	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	log.Printf("Sending Hello to Client")
	return &pb.HelloResponse{
		Message: "Hello from Server - unary.go",
	}, nil
}
