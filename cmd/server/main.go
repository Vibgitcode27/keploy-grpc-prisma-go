package main

import (
	"context"
	"flag"
	"fmt"
	"krpg/krpg"
	"krpg/service"
	"log"
	"net"

	"github.com/keploy/go-sdk/keploy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "The port to start the server on")
	flag.Parse()
	log.Printf("Starting server on port %d", *port)

	k := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "my-grpc-app",
			Port: fmt.Sprintf("%d", *port),
		},
		Server: keploy.ServerConfig{
			URL: "http://localhost:6789/api",
		},
	})

	todoServer := service.NewTodoServer()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(yourUnaryInterceptor(k)),
	)

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

func yourUnaryInterceptor(k *keploy.Keploy) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		k.Log.Info("Received request", zap.String("method", info.FullMethod))

		resp, err = handler(ctx, req)

		if err != nil {
			k.Log.Error("Error processing request", zap.String("method", info.FullMethod), zap.Error(err))
		}

		return resp, err
	}
}
