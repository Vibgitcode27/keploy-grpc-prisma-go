package service

import (
	"context"
	"krpg/krpg"
	"log"

	"github.com/google/uuid"
)

type todoServer struct {
}

func NewTodoServer() *todoServer {
	return &todoServer{}
}

func (server *todoServer) Create(ctx context.Context, req *krpg.CreateRequest) (res *krpg.CreateResponse, err error) {

	log.Print("Create request received with title: ", req.Title)

	newTodo := krpg.Todo{
		Id:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		DueDate:     "2021-09-01",
		Completed:   false,
	}
	log.Print("New todo created: \n", newTodo)
	return &krpg.CreateResponse{Task: &newTodo}, nil
}
