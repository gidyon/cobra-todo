package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"

	"github.com/gidyon/ctodo/todo"
)

func main() {
	srv := grpc.NewServer()
	var todoServer *todoTaskManager
	todo.RegisterTodoManagerServer(srv, todoServer)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("could not listen on tcp :9090: %v\n", err)
	}
	log.Println("Server started on port :9090, waiting for rpc calls...")
	log.Fatalln(srv.Serve(l))
}

const (
	dbPath = "mytodo.pb"
)

func readTodos() (*todo.Todo, error) {
	var todolist todo.Todo

	f, err := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("could not open file %v: %v", dbPath, err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read file %v: %v", dbPath, err)
	}
	if len(b) == 0 {
		todolist.Tasks = make([]*todo.Task, 0)
		return &todolist, nil
	}

	if err := proto.Unmarshal(b, &todolist); err != nil {
		return nil, fmt.Errorf("could not read tasks: %v", err)
	}
	return &todolist, nil
}

type todoTaskManager struct{}

func (tsv *todoTaskManager) List(context.Context, *todo.Void) (*todo.Todo, error) {
	tasks, err := readTodos()
	if err != nil {
		return nil, fmt.Errorf("reading todos failed: %v", err)
	}
	return tasks, nil
}

func (tsv *todoTaskManager) Add(ctx context.Context, task *todo.Task) (*todo.Void, error) {
	tasks, err := readTodos()
	if err != nil {
		return nil, fmt.Errorf("reading todos failed: %v", err)
	}

	tasks.Tasks = append(tasks.Tasks, task)

	err = marshalAndwriteToFile(tasks)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not marshal and write to file: %v", err)
	}

	return &todo.Void{}, nil
}

func (tsv *todoTaskManager) Delete(ctx context.Context, taskID *todo.TaskId) (*todo.Void, error) {
	tasks, err := readTodos()
	if err != nil {
		return nil, fmt.Errorf("reading todos failed: %v", err)
	}

	if len(tasks.Tasks) == 0 {
		return nil, fmt.Errorf("no item in todo, please add items first")
	}

	for i, task := range tasks.Tasks {
		if task.GetTaskId() == taskID.GetTaskId() {
			tasks.Tasks = append(tasks.Tasks[i:], tasks.Tasks[i+2:]...)
			break
		}
	}

	err = marshalAndwriteToFile(tasks)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not marshal and write to file: %v", err)
	}

	return &todo.Void{}, nil
}

func (tsv *todoTaskManager) Done(ctx context.Context, taskID *todo.TaskId) (*todo.Void, error) {
	tasks, err := readTodos()
	if err != nil {
		return nil, fmt.Errorf("reading todos failed: %v", err)
	}

	if len(tasks.Tasks) == 0 {
		return nil, fmt.Errorf("no item in todo, please add items first")
	}

	for _, task := range tasks.Tasks {
		if task.GetTaskId() == taskID.GetTaskId() {
			task.Done = true
			break
		}
	}

	err = marshalAndwriteToFile(tasks)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not marshal and write to file: %v", err)
	}

	return &todo.Void{}, nil
}

func marshalAndwriteToFile(msg proto.Message) error {
	b, err := proto.Marshal(msg)
	if err != nil {
		return status.Errorf(codes.Internal, "could not encode back tasks: %v", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("could not open file %v: %v", dbPath, err)
	}

	if _, err = f.Write(b); err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}

	if err = f.Close(); err != nil {
		return fmt.Errorf("could not close file %v %v", dbPath, err)
	}
	return nil
}
