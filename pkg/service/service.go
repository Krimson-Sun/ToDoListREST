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
	UpdateList(userId, listId int, input todo.UpdateTodo) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, lisId int) ([]todo.TodoItem, error)
	GetById(userId, ItemId int) (todo.TodoItem, error)
	Update(userId, itemId int, updateInput todo.UpdateItem) error
	Delete(userId, itemId int) error
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
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
