package main

import (
	"context"
	"flag"
	"fmt"
	"krpg/krpg"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "The server address in the format of host:port")
	flag.Parse()
	fmt.Printf("Dial server on address %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatal("Error dialing server: ", err)
	}

	todoClient := krpg.NewTodoServiceClient(conn)
	req := &krpg.CreateRequest{
		Title:       "Learn gRPC",
		Description: "Learn how to use gRPC in Go",
		DueDate:     "2021-09-01",
	}

	res, err := todoClient.Create(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			fmt.Println("Laptop already exists")
		} else {
			log.Fatal("Cannot create laptop(Response)", err)
		}
		return
	}

	fmt.Printf("Todo created with id: %s", res.Task.Id)

}
