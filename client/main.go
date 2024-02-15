package main

import (
	"log"

	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
	//Both need to listen on same port
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to Connect the Grpc server: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	callSayHello(client)

}
