package service

import (
	todo "TodoListREST"
	"TodoListREST/pkg/repository"
)

type TodolistService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodolistService {
	return &TodolistService{repo: repo}
}

func (s *TodolistService) Create(userId int, list todo.Todo) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodolistService) GetAll(userId int) ([]todo.Todo, error) {
	return s.repo.GetAll(userId)
}

func (s *TodolistService) GetById(userId, ListId int) (todo.Todo, error) {
	return s.repo.GetById(userId, ListId)
}

func (s *TodolistService) DeleteList(userId, listId int) error {
	return s.repo.DeleteList(userId, listId)
}

func (s *TodolistService) UpdateList(userId, listId int, input todo.UpdateTodo) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(userId, listId, input)
}
