package repository

import (
	todo "TodoListREST"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.Todo) (int, error)
	GetAll(userId int) ([]todo.Todo, error)
	GetById(userId, listId int) (todo.Todo, error)
	DeleteList(userId, listId int) error
	UpdateList(userId, listId int, input todo.UpdateTodo) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
