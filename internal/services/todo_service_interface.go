package services

import (
	"testrestapi/internal/models"
)

type TodoService interface {
	GetTodos() ([]models.Todo, error)
	GetTodoByID(id int) (models.Todo, error)
	AddTodo(todo models.Todo) (models.Todo, error)
	ToggleTodoStatus(id int) (models.Todo, error)
}
