package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start the GRPC server: %v", err)
	}
}
