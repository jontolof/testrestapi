package services

import (
	"testrestapi/internal/models"
	"testrestapi/internal/repository"
)

type todoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) TodoService {
	return &todoService{repository: repository}
}

// MARK: - Handle GET
func (service *todoService) GetTodos() ([]models.Todo, error) {
	return service.repository.GetAll()
}

func (service *todoService) GetTodoByID(id int) (models.Todo, error) {
	return service.repository.GetByID(id)
}

// MARK: - Handle POST
func (service *todoService) AddTodo(todo models.Todo) (models.Todo, error) {
	return service.repository.Create(todo)
}

// MARK: - Handle PATCH
func (service *todoService) ToggleTodoStatus(id int) (models.Todo, error) {
	todo, err := service.repository.GetByID(id)

	if err != nil {
		return models.Todo{}, err
	}

	todo.Completed = !todo.Completed
	return service.repository.Update(todo)
}
