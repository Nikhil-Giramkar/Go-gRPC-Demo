package main

import (
	"log"
	"net"

	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("Server started at: %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start the GRPC server: %v", err)
	}
}
