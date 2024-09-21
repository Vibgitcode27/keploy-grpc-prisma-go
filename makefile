gen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

server:
	go run cmd/server/main.go -port 8000

client:
	go run cmd/client/main.go -address 0.0.0.0:8000

keploy:
	GOPATH=$(HOME)/go keploy record -c "/usr/local/go/bin/go run cmd/server/main.go -port 8000"

start:
	sudo -E make keploy