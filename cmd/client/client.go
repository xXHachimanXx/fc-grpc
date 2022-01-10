package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/xXHachimanXx/fc-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Juca",
		Email: "jooj@jjj.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Juca",
		Email: "jooj@jjj.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status: ", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Felipe",
			Email: "felipe@email.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Felipe 2",
			Email: "felipe2@email.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Felipe3",
			Email: "felipe3@email.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Felipe4",
			Email: "felipe4@email.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Felipe5",
			Email: "felipe5@email.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
