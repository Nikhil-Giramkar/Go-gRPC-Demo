package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
)

func callSayHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bi-Directional Streaming started")

	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Error in sending request: %v", err)
	}

	//For sending and receiving paralley via streams
	//Lets use go routine for receiving and wait for that routine using channel
	myChan := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error in receiving response: %v", err)

			}
			log.Printf(message.Message)
		}

		close(myChan)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to Send stream : %v", err)
		}

		// 2 second delay to simulate a long running process
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-myChan
	log.Printf("Bi-Directional Streaming Finished")
}
