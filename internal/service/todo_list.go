package service

import (
	"github.com/rogudator/todo-app/entity"
	"github.com/rogudator/todo-app/internal/repository"
)

type TodoList interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input entity.UpdateListInput) error
}

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list entity.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]entity.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (entity.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s* TodoListService) Update(userId, listId int, input entity.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	
	return s.repo.Update(userId, listId, input)
}