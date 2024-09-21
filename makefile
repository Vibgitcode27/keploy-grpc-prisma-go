gen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go
