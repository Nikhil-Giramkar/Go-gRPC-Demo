package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Failed to read from SayHelloServerStreaming(): %v", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF { //End of file
			break
		}

		if err != nil {
			log.Fatalf("Error on reading stream: %v ", err)
		}

		log.Printf(message.Message)
	}

	log.Printf("Streaming Finished")
}
