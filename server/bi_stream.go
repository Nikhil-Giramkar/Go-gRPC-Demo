package main

import (
	"io"
	"log"

	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
)

func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error in receiving request: %v", err)
			return err
		}
		log.Printf("Got request with name : %v", req.Name)

		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		if err := stream.Send(res); err != nil {
			log.Fatalf("Error in sending response: %v", err)
		}

	}
}
