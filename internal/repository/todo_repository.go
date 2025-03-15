package repository

import (
	"database/sql"
	"testrestapi/internal/models"
)

type todoRepository struct {
	DB *sql.DB
}

// NewTodoRepository returns a new instance of TodoRepository with a given *sql.DB
func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{DB: db}
}

func (repository *todoRepository) GetAll() ([]models.Todo, error) {
	rows, err := repository.DB.Query("SELECT id, item, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Item, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (repository *todoRepository) GetByID(id int) (models.Todo, error) {
	var todo models.Todo
	err := repository.DB.QueryRow("SELECT id, item, completed FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Item, &todo.Completed)
	return todo, err
}

func (repository *todoRepository) Create(todo models.Todo) (models.Todo, error) {
	err := repository.DB.QueryRow(
		"INSERT INTO todos (item, completed) VALUES ($1, $2) RETURNING id",
		todo.Item, todo.Completed,
	).Scan(&todo.ID)
	return todo, err
}

func (repository *todoRepository) Update(todo models.Todo) (models.Todo, error) {
	_, err := repository.DB.Exec("UPDATE todos SET item = $1, completed = $2 WHERE id = $3",
		todo.Item, todo.Completed, todo.ID)
	return todo, err
}
