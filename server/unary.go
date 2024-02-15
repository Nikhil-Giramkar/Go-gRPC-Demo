package main

import (
	"context"
	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from Server - unary.go",
	}, nil
}
