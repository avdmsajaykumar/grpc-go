package main

import (
	"context"
	"fmt"
	greetpb "grpc-go-course/greet1/proto-files"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var mu sync.Mutex

func main() {
	// filedata, err := ioutil.ReadFile("C:\\Users\\avdms\\go\\src\\grpc-go-course\\greet1\\client\\clientdata.txt")
	// data := strings.Split(string(filedata), ",")
	// fn, err := strconv.Atoi(data[0])
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// ln, err := strconv.Atoi(data[1])
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fn32 := int32(fn)
	// ln32 := int32(ln)

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error Establishing Connection :%v\n", err)
	}
	defer con.Close()
	// wholeChannel1 := make(chan int)
	// wholeChannel2 := make(chan int)
	// wholeChannel3 := make(chan int)
	// wholeChannel4 := make(chan int)

	client := greetpb.NewGreetServiceClient(con)

	// go doUnary(client, fn32, ln32, wholeChannel1, wholeChannel2)
	// go doServerStreaming(client, wholeChannel2, wholeChannel3)
	// go doClientStreaming(client, wholeChannel3, wholeChannel4)
	// go doBidirectionalStreaming(client, wholeChannel4, wholeChannel1)
	// doSquareRoot(client, 10)
	// doSquareRoot(client, 81)
	// doSquareRoot(client, -4)
	doDeadlineUnary(client, 1)
	doDeadlineUnary(client, 3)
	doDeadlineUnary(client, 5)

	// <-wholeChannel1
	// //close(wholeCHannel)

	// x := <-wholeCHannel
	// y := <-wholeCHannel
	// z := <-wholeCHannel
	// xyz := <-wholeCHannel
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(xyz)
	// go func() {
	// 	time.Sleep(20 * time.Second)
	// 	close(wholeCHannel)
	// }()

	// <-wholeCHannel

}

func doUnary(client greetpb.GreetServiceClient, fn, ln int32, w1 <-chan int, w2 chan<- int) {
	// <-w1
	in := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstNumber: fn,
			LastNumber:  ln,
		},
	}
	fmt.Println("Calling unary")
	Response, error := client.GreetUnary(context.Background(), in)
	if error != nil {
		log.Fatalf("Error at Calling  :%v\n", error)
		// close(w1)

	}
	fmt.Println("Unary: ", Response.GetResult())

	fmt.Println("Unary: It's the End")
	w2 <- 1
}

func doServerStreaming(client greetpb.GreetServiceClient, w1 <-chan int, w2 chan<- int) {
	<-w1
	in := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstNumber: 30,
		},
	}
	Res, Err := client.GreetServerStreaming(context.Background(), in)
	if Err != nil {
		log.Fatalf("%v", Err)
		// close(w1)
	}

	for {
		msg, err := Res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
			// close(w1)
		}
		fmt.Println(msg.GetResult())
	}
	w2 <- 2

}

func doClientStreaming(client greetpb.GreetServiceClient, w1 <-chan int, w2 chan<- int) {
	<-w1

	number := []int{3, 4, 5, 6, 7, 8, 9}

	stream, error := client.ClientStreaming(context.Background())
	if error != nil {
		log.Fatalln(error)
		// close(w1)
	}

	for _, indx := range number {
		fmt.Println(indx)
		stream.Send(&greetpb.AverageRequest{
			Number: float32(indx),
		})

	}
	Resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
		// close(w1)
	}
	fmt.Println("Client S Result is ", Resp.GetAverage())
	w2 <- 3
}

func doBidirectionalStreaming(client greetpb.GreetServiceClient, w1 <-chan int, w2 chan<- int) {
	<-w1
	requests := []*greetpb.Greet1Request{
		&greetpb.Greet1Request{
			FirstName: "Ajay",
			LastName:  "Kumar",
		},
		&greetpb.Greet1Request{
			FirstName: "Venkata",
			LastName:  "Durga",
		},
		&greetpb.Greet1Request{
			FirstName: "Maruthi",
			LastName:  "Shanmuka",
		},
	}

	stream, err := client.BideStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while invoking Client stream: %v", err)
	}
	waitchannel := make(chan struct{})
	go func() {
		for _, request := range requests {
			fmt.Printf("Sending Data: %s\n", request)
			error := stream.Send(request)
			if error != nil {
				close(waitchannel)

			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving Server Response stream: %v", err)
				// close(w1)
				break
			}
			fmt.Printf("%s", resp.GetResult())
		}
		close(waitchannel)

	}()
	<-waitchannel
	w2 <- 4
}

func doSquareRoot(client greetpb.GreetServiceClient, NewNumber int32) {
	fmt.Println("Invoked squareroot funtion")
	// NewNumber := int32(10)
	response, error := client.SquareRoot(context.Background(), &greetpb.SquareRootRequest{Number: NewNumber})
	if error != nil {
		receivedErr, ok := status.FromError(error)
		if ok {
			fmt.Println(receivedErr.Code())
			if receivedErr.Code() == codes.InvalidArgument {
				fmt.Println("We Sent a Negative Number probably")
				return
			}
		} else {
			log.Fatalf("Received a system Exception %v", receivedErr)
		}
	}
	fmt.Println(response.GetNumber())
}

func doDeadlineUnary(client greetpb.GreetServiceClient, sec time.Duration) {
	fmt.Println("Entered Deadline Method")

	in := &greetpb.GreetWithDeadlineRequest{
		FirstName: "Ajay Kumar",
	}

	ctx, Cancel := context.WithTimeout(context.Background(), sec*time.Second)

	defer Cancel()

	response, error := client.GreetWithDeadline(ctx, in)
	if error != nil {
		respErr, ok := status.FromError(error)
		if ok {
			if respErr.Code() == codes.DeadlineExceeded {
				fmt.Println("DeadLine Exceeded")
			} else {
				log.Fatalf("Unexcepted Error %v", respErr)
			}

		} else {
			log.Fatalf("Unexcepted Error calling grpc %v", respErr)
		}
	}

	fmt.Println(response.GetResult())
}
