gen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

server:
	go run cmd/server/main.go -port 8000

client:
	go run cmd/client/main.go -address 0.0.0.0:8000

keploy:
	keploy record --command "/usr/local/go/bin/go run cmd/server/main.go -port 8000" --proxy-port 9000

start:
	sudo -E make keploy
