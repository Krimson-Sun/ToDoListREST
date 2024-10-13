package service

import (
	todo "TodoListREST"
	"TodoListREST/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.Todo) (int, error)
	GetAll(userId int) ([]todo.Todo, error)
	GetById(userId, listId int) (todo.Todo, error)
	DeleteList(userId, listId int) error
	UpdateList(usesrId, listId int, input todo.UpdateTodo) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
