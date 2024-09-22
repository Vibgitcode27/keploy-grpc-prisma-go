package main

import (
	"flag"
	"fmt"
	"krpg/krpg"
	"krpg/service"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/keploy/go-sdk/integrations/kgrpcserver"
	"github.com/keploy/go-sdk/keploy"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting gRPC server with Keploy")

	port := flag.Int("port", 8000, "The port to start the server on")
	flag.Parse()
	fmt.Printf("Parsed port: %d\n", *port)

	k := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "my-grpc-app",
			Port: fmt.Sprintf("%d", *port),
		},
	})
	fmt.Println("Keploy instance created")

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

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			fmt.Printf("Failed to serve: %v\n", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Printf("Received signal: %v. Shutting down server...\n", sig)

	grpcServer.GracefulStop()
	fmt.Println("gRPC server stopped gracefully")
}
