package main

import (
	"flag"
	"fmt"
	"krpg/krpg"
	"krpg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// We get port from command line arguments
	port := flag.Int("port", 0, "The port to start the server on")
	flag.Parse()
	log.Printf("Starting server on port %d", *port)

	todoServer := service.NewTodoServer()
	grpcServer := grpc.NewServer()

	krpg.RegisterTodoServiceServer(grpcServer, todoServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
