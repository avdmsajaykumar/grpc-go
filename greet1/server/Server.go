package main

import (
	"context"
	"fmt"
	greetpb "grpc-go-course/greet1/proto-files"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SquareRoot(ctx context.Context, req *greetpb.SquareRootRequest) (*greetpb.SquareRootResponse, error) {
	fmt.Println("Received a SquareRoot Request")
	number := req.GetNumber()
	if number < 0 {
		error := status.Error(codes.InvalidArgument, fmt.Sprintf("Received a negative number Which is Invalid"))
		fmt.Println("Received a invalid Request")
		return nil, error
	}

	result := math.Sqrt(float64(number))
	response := &greetpb.SquareRootResponse{Number: result}
	fmt.Println("Sent Response")
	return response, nil
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Println("Request received")
	request := req.GetFirstName()
	Result := "Hello " + request + " !!!"

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		if ctx.Err() != nil {
			if ctx.Err() == context.Canceled {
				fmt.Println("Client Cancelled the request")
				return nil, status.Error(codes.Canceled, "Client Cancelled the request")
			}
		}
	}
	res := &greetpb.GreetWithDeadlineResponse{
		Result: Result,
	}

	return res, nil
}

func (*server) GreetUnary(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("received a Request")
	firstNumber := req.GetGreeting().GetFirstNumber()
	lastNumber := req.GetGreeting().GetLastNumber()
	fmt.Printf("Numbers received are %d, %d\n", firstNumber, lastNumber)
	res := firstNumber + lastNumber

	response := &greetpb.GreetResponse{
		Result: res,
	}

	return response, nil
}

func (*server) GreetServerStreaming(req *greetpb.GreetRequest, stream greetpb.GreetService_GreetServerStreamingServer) error {
	firstname := req.GetGreeting().GetFirstNumber()
	for i := 0; i <= 10; i++ {
		res := &greetpb.GreetResponse{
			Result: firstname + int32(i),
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)

	}
	return nil

}
func (*server) ClientStreaming(stream greetpb.GreetService_ClientStreamingServer) error {
	fmt.Println("Enter Client Streaming")
	sum := float32(0)
	count := float32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := sum / count
			stream.SendAndClose(&greetpb.AverageResponse{
				Average: float64(result),
			})
			return nil

		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(req.GetNumber())
		sum += req.GetNumber()
		count++

	}
}

func (*server) BideStreaming(stream greetpb.GreetService_BideStreamingServer) error {
	fmt.Println("Invoked Biderectional Streaming")
	result := ""
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Reached EOF")
			return nil
		}
		if err != nil {
			log.Fatalf("error reading from client %v", err)
		}

		Firstname := req.GetFirstName()
		Lastname := req.GetLastName()
		result = "Hello " + Firstname + " " + Lastname + " !\n"
		error := stream.Send(&greetpb.Greet1Response{
			Result: result,
		})
		if error != nil {
			return error
		}

	}
}

func main() {
	fmt.Println("Started the Serer")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error Listening :%v\n", err)
	}
	defer lis.Close()
	grpcserver := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(grpcserver, &server{})

	if err := grpcserver.Serve(lis); err != nil {
		log.Fatalf("Error Serving :%v\n", err)
	}
}
