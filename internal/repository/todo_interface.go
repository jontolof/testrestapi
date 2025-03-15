package repository

import (
	"testrestapi/internal/models"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetById(id int) (models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
}
