package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Nikhil-Giramkar/Go-gRPC-Demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParams{})

	if err != nil {
		log.Fatalf("Failed to read from SayHello(): %v", err)
	}

	log.Printf("%s", res.Message)
}
