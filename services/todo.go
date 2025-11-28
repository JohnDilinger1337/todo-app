package services

import (
	"fmt"
	"main/models"
	"strconv"
)

type TodoService interface {
	Add(input models.TodoInput) error
	Delete(input int) error
	GetAll() []models.Todo
}

type todoService struct {
	todos  []models.Todo
	nextID int
}

func NewTodoService() TodoService {
	return &todoService{
		todos:  []models.Todo{},
		nextID: 1,
	}
}

func (s *todoService) GetAll() []models.Todo {
	return s.todos
}

func (s *todoService) Add(input models.TodoInput) error {
	if input.Title == "" {
		return fmt.Errorf("invalid title: %s", input.Title)
	} else if !input.Category.IsValid() {
		return fmt.Errorf("invalid category: %s", input.Category)
	}

	todo := models.Todo{
		ID:       s.nextID,
		Title:    input.Title,
		Category: input.Category,
		Done:     false,
	}
	s.todos = append(s.todos, todo)
	s.nextID++

	return nil
}

func (s *todoService) Delete(input int) error {
	if s.todos == nil || input <= 0 || input > len(s.todos) {
		return fmt.Errorf("something went wrong while deleteing this item with ID: (%s)", strconv.Itoa(input))
	}

	s.todos = append((s.todos)[:input-1], (s.todos)[input:]...)

	return nil
}
