package main

import (
	"flag"
	"fmt"
	"krpg/krpg"
	"krpg/service"
	"net"
	"os"

	"github.com/keploy/go-sdk/integrations/kgrpcserver"
	"github.com/keploy/go-sdk/keploy"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting main function")
	port := flag.Int("port", 0, "The port to start the server on")
	flag.Parse()
	fmt.Printf("Parsed port: %d\n", *port)

	k := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "my-grpc-app",
			Port: fmt.Sprintf("%d", *port),
		},
		Server: keploy.ServerConfig{
			URL: "http://localhost:6789/api",
		},
	})
	fmt.Println("Keploy instance created")

	fmt.Println("Starting gRPC server setup")
	todoServer := service.NewTodoServer()
	grpcServer := grpc.NewServer(kgrpcserver.UnaryInterceptor(k))
	krpg.RegisterTodoServiceServer(grpcServer, todoServer)
	fmt.Println("gRPC server setup completed")

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	fmt.Printf("Attempting to listen on: %s\n", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Listener created successfully")

	fmt.Println("Starting gRPC server")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
