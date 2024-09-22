FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main", "-port", "8000"]