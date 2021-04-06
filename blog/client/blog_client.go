package main

import (
	"context"
	"fmt"
	blogpb "grpc-go-course/blog/protofiles"
	"log"

	"google.golang.org/grpc"
)

func main() {

	fmt.Print("Started Client ")
	grpcConnect, ConnecterError := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if ConnecterError != nil {
		log.Fatalf("Error obtaining Connection: %v \n", ConnecterError)
	}

	grpcClient := blogpb.NewBlogServiceClient(grpcConnect)

	//insertOne(grpcClient)
	//ReadOne(grpcClient, "5e7876fe5c32207ed4dbb7a6")
	//ReadOne(grpcClient, "5e787a15129580409d7026b2")
	// ReadOne(grpcClient, "5e787a15129580409d7066b2")
	//UpdateOne(grpcClient, "5e7877d7e9d7dbaf746333ef")
	deleteOne(grpcClient, "5e7877d7e9d7af746433ef")

}

func insertOne(client blogpb.BlogServiceClient) {
	fmt.Println("Invoked A Insert Function")

	input := &blogpb.CreateEntryRequest{
		EntryRequest: &blogpb.EntryRequest{
			Name: "JAVA",
			Type: "Static",
			Year: "1995",
		},
	}

	Response, ResponseError := client.CreateEntry(context.Background(), input)
	if ResponseError != nil {
		log.Fatalf("Error While receiving Response : %v ", ResponseError)
	}
	fmt.Printf("\n Generated DB ID is %v \n", Response.GetId())

}

func ReadOne(Client blogpb.BlogServiceClient, id string) {
	fmt.Println("invoked Read Call with : " + id)
	request := blogpb.ReadEntryRequest{
		Id: id,
	}

	Response, RespErr := Client.ReadEntry(context.Background(), &request)
	if RespErr != nil {
		log.Fatalf("Error in response : %v", RespErr)
	}
	fmt.Printf("Read Entry Result is %v \n\n", Response)

}

func UpdateOne(client blogpb.BlogServiceClient, id string) {
	fmt.Println("invoked Update Call with : " + id)

	request := &blogpb.UpdateEntryRequest{
		EntryRequest: &blogpb.EntryRequest{
			Id:   id,
			Name: "Go",
			Type: "STATIC",
			Year: "2009",
		},
	}

	response, resErr := client.UpdateEntry(context.Background(), request)
	if resErr != nil {
		log.Fatalf("Erroring while Updating :%v \n", resErr)
	}
	fmt.Printf("Response Status : %s \n", response.GetStatus())

}

func deleteOne(client blogpb.BlogServiceClient, str string) {
	fmt.Println("Deleting a Document from DB")
	response, err := client.DeleteEntry(context.Background(), &blogpb.DeleteEntryRequest{
		Id: str,
	})
	if err != nil {
		log.Fatalf("Error while Deleting : %v \n", err)
	}

	fmt.Printf("Result Status :%v \n", response.GetResult())
}
