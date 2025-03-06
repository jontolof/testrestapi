package services

import (
	"errors"
	"sync"
	"testrestapi/internal/models"
)

type TodoService struct {
	todos []models.Todo
	mu    sync.Mutex
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos: []models.Todo{
			{ID: "1", Item: "Clean Room", Completed: false},
			{ID: "2", Item: "Read Book", Completed: false},
			{ID: "3", Item: "Record Video", Completed: false},
		},
	}
}

// MARK: - Handle GET
func (service *TodoService) GetTodos() ([]models.Todo, error) {
	return service.todos, nil
}

func (service *TodoService) GetTodoById(id string) (*models.Todo, error) {
	for i, t := range service.todos {
		if t.ID == id {
			return &service.todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

// MARK: - Handle POST
func (service *TodoService) AddTodo(todo *models.Todo) (*models.Todo, error) {
	service.todos = append(service.todos, *todo)
	return todo, nil
}

// MARK: - Handle PATCH

func (service *TodoService) ToggleTodoStatus(id string) (*models.Todo, error) {
	todo, err := service.GetTodoById(id)

	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed
	return todo, nil
}
