package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client Streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Failed to read from SayHelloServerStreaming(): %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to Send stream : %v", err)
		}

		log.Printf("Sent request with name : %s", name)
		// 2 second delay to simulate a long running process
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	log.Printf(" Client Streaming Finished")

	if err != nil {
		log.Fatalf("Failed to Receive response : %v", err)
	}

	log.Printf("Response: %v", res.Messages)
}
