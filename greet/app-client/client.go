package main

import (
	"fmt"
	"log"

	greetpb "grpc-go-course/greet/proto-files"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client is running")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect %v\n", err)
	}

	client := greetpb.NewGreatServiceClient(conn)
	fmt.Println(client)

}
