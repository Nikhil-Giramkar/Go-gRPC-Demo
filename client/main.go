package main

import (
	"fmt"
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

	var choice int
	nameList := &pb.NameList{
		Names: []string{"Nikhil", "Ansh", "Maria"},
	}

	for {
		fmt.Println("What type of gRPC Communication you wish to see - ")
		fmt.Println("1. Request-Response (Unary) ")
		fmt.Println("2. Server Streamimg ")
		fmt.Println("3. Client Streaming")
		fmt.Println("4. Bi-Directional Streaming ")

		_, error := fmt.Scanln(&choice)
		if error != nil {
			fmt.Println("Try again, error in taking input")
			break
		}
		print("-------------------------------------------------------------------\n")

		switch choice {
		case 1:
			//Unary
			callSayHello(client)
		case 2:
			//ServerStream
			callSayHelloServerStream(client, nameList)
		case 3:
			//Client Stream
			callSayHelloClientStream(client, nameList)
		case 4:
			//Bi-Directional
			callSayHelloBiDirectionalStream(client, nameList)
		default:
			fmt.Println("Try Again")
			break
		}
		print("-------------------------------------------------------------------\n")
	}

}
