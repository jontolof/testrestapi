package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"testrestapi/internal/models"
	"testrestapi/internal/services"
)

type TodoController struct {
	Service *services.TodoService
}

// Create an instance of TodoController
func NewTodoController(service *services.TodoService) *TodoController {
	return &TodoController{
		Service: service,
	}
}

// MARK: - Handle GET
func (controller *TodoController) GetTodos(responseWriter http.ResponseWriter, request *http.Request) {
	todos, err := controller.Service.GetTodos()
	if err != nil {
		http.Error(responseWriter, "Bad request", http.StatusBadRequest)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(todos)
}

func (controller *TodoController) GetTodo(responseWriter http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.Path, "/")

	if len(parts) < 3 || parts[2] == "" {
		http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
		return
	}

	id := parts[2]
	todo, err := controller.Service.GetTodoById(id)

	if err != nil {
		http.Error(responseWriter, "Todo not found", http.StatusNotFound)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(todo)
}

// MARK: - Handle POST
func (controller *TodoController) AddTodo(responseWriter http.ResponseWriter, request *http.Request) {
	var newTodo models.Todo

	if err := json.NewDecoder(request.Body).Decode(&newTodo); err != nil {
		http.Error(responseWriter, "Bad request", http.StatusBadRequest)
		return
	}

	todo, err := controller.Service.AddTodo(&newTodo)

	if err != nil {
		http.Error(responseWriter, "Bad request", http.StatusBadRequest)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(todo)
}

// MARK: - Handle PATCH
func (controller *TodoController) ToggleTodoStatus(responseWriter http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := parts[2]

	todo, err := controller.Service.ToggleTodoStatus(id)

	if err != nil {
		http.Error(responseWriter, "Todo not found", http.StatusNotFound)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(todo)
}
