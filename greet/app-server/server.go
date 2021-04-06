package main

import (
	"fmt"
	greatpb "grpc-go-course/greet/proto-files"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Server is running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error while listening%v\n", err)
	}

	s := grpc.NewServer()

	greatpb.RegisterGreatServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error while Serving%v\n", err)
	}

}
