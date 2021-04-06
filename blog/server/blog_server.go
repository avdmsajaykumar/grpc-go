package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	blogpb "grpc-go-course/blog/protofiles"
)

type ServerStruct struct{}

type Items struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"Name"`
	CompilerType string             `bson:"Compiler"`
	Year         string             `bson:"Year"`
}

var mongocollection *mongo.Collection

func (*ServerStruct) CreateEntry(ctx context.Context, in *blogpb.CreateEntryRequest) (*blogpb.ResponseCreateEntry, error) {
	response := &blogpb.ResponseCreateEntry{}
	request := in.GetEntryRequest()
	var data Items
	if request != nil {
		data.Name = request.GetName()
		data.CompilerType = request.GetType()
		data.Year = request.GetYear()

	} else {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Couldn't find Request : %v", request))
	}

	mongoctx, mongoCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer mongoCancel()
	insertResult, insertError := mongocollection.InsertOne(mongoctx, &data)
	if insertError != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", insertError),
		)
	}
	response.Id = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return response, nil

}

func (*ServerStruct) ReadEntry(ctx context.Context, req *blogpb.ReadEntryRequest) (*blogpb.ReadEntryResponse, error) {
	var data Items

	id := req.GetId()
	objectid, ok := primitive.ObjectIDFromHex(id)
	if ok != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID passed is not a Hex String : %v", ok))
	}
	filter := bson.M{
		"_id": objectid,
	}

	mangoResult := mongocollection.FindOne(context.Background(), filter)

	Decodeerror := mangoResult.Decode(&data)
	if Decodeerror != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ERROR : %v ", Decodeerror))
	}

	return &blogpb.ReadEntryResponse{EntryRequest: &blogpb.EntryRequest{
		Name: data.Name,
		Year: data.Year,
		Type: data.CompilerType,
		Id:   data.Id.Hex(),
	}}, nil

}
func (*ServerStruct) UpdateEntry(ctx context.Context, req *blogpb.UpdateEntryRequest) (*blogpb.UpdateEntryResponse, error) {
	id := req.GetEntryRequest().GetId()
	ObID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Invalid Id : %v", id))
	}

	filter := bson.M{
		"_id": ObID,
	}

	mangoResult := mongocollection.FindOne(context.Background(), filter)

	if err := mangoResult.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Couldn't find the entry : %v\n", err))
		}
	}
	fmt.Println("Data Retrive Success")

	data := Items{}

	if req.GetEntryRequest() != nil {
		data.Name = req.GetEntryRequest().GetName()
		data.CompilerType = req.GetEntryRequest().GetType()
		data.Year = req.GetEntryRequest().GetYear()

		updateResult, updateErr := mongocollection.ReplaceOne(context.Background(), filter, &data)
		if updateErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("ERROR : %v", updateErr))
		}

		filtercount := updateResult.MatchedCount
		updatedcount := updateResult.ModifiedCount

		fmt.Printf("No of Documents obtained : %v \n", filtercount)
		fmt.Printf("No of Documents updated : %v \n", updatedcount)
		// fmt.Printf("OID/new OID : %v \n", updatedID)
		return &blogpb.UpdateEntryResponse{Status: "SUCCUSS"}, nil
	} else {
		return &blogpb.UpdateEntryResponse{Status: "Fail"}, nil
	}

}

func (*ServerStruct) DeleteEntry(ctx context.Context, req *blogpb.DeleteEntryRequest) (*blogpb.DeleteEntryResponse, error) {
	id := req.GetId()
	objectid, OIDerror := primitive.ObjectIDFromHex(id)
	if OIDerror != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Couldn't convert string : %v\n ", OIDerror))
	}

	filter := bson.M{
		"_id": objectid,
	}

	deleteResponse, deleteErr := mongocollection.DeleteOne(context.Background(), filter)
	if deleteErr != nil {
		return nil, deleteErr
	}
	fmt.Printf("%v Entry is deleted \n  ", deleteResponse.DeletedCount)
	if deleteResponse != nil {
		return &blogpb.DeleteEntryResponse{Result: "Success"}, nil
	} else {
		return &blogpb.DeleteEntryResponse{Result: "Failed"}, nil
	}
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Started Server")
	//Creating a Server Listner
	lis, liserr := net.Listen("tcp", "localhost:50051")
	if liserr != nil {
		log.Fatalf("could enable the listener : %v  \n", liserr)
	}

	defer lis.Close()
	//Creating a gRPC Server
	grpcServer := grpc.NewServer()

	//Registering gRPC server to Protocal Buffer
	blogpb.RegisterBlogServiceServer(grpcServer, &ServerStruct{})

	fmt.Println("Creating Mongo Client")
	//Connected to a mongoServer
	mongoclient, mongoClienterr := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if mongoClienterr != nil {
		log.Fatalf("Error Creating MongoClient: %v", mongoClienterr)
	}

	//Creating a Timeout Context for Mongo DB
	mongoctx, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoCancel()

	//Connecting to Mongo Client
	fmt.Println("Connecting to MongoDB")
	mongoConnecterr := mongoclient.Connect(mongoctx)
	if mongoConnecterr != nil {
		log.Fatalf("Error While Connecting to Mongo: %v \n", mongoConnecterr)
	}

	mongocollection = mongoclient.Database("DB1").Collection("Programming_Languages")

	//Started listening

	go func() {
		fmt.Println("Started Listening")
		SerLisError := grpcServer.Serve(lis)
		if SerLisError != nil {
			log.Fatalf("Unable to Serve : %v \n", SerLisError)
		}
	}()

	//Creating a Interrupt Channel
	InterruptChannel := make(chan os.Signal, 1)
	// Waiting to receive the interrupt signal
	signal.Notify(InterruptChannel, os.Interrupt)

	<-InterruptChannel

	fmt.Println("Received the Interrupt Signal....\nClosing Connections")

	fmt.Println("Closing Mongo Connection")
	mongoclient.Disconnect(context.TODO())

	fmt.Println("Stopping gRPC Server")
	grpcServer.Stop()

}
