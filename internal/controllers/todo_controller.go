package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"testrestapi/internal/models"
	"testrestapi/internal/services"
)

type TodoController struct {
	Service services.TodoService
}

// Create an instance of TodoController
func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{Service: service}
}

// MARK: - Handle GET
func (controller *TodoController) GetTodos(writer http.ResponseWriter, request *http.Request) {
	todos, err := controller.Service.GetTodos()
	if err != nil {
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(todos)
}

func (controller *TodoController) GetTodo(writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.Path, "/")

	if len(parts) < 3 || parts[2] == "" {
		http.Error(writer, "Missing ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(writer, "Unvalid ID", http.StatusBadRequest)
		return
	}
	todo, err := controller.Service.GetTodoByID(id)

	if err != nil {
		http.Error(writer, "Todo not found", http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(todo)
}

// MARK: - Handle POST
func (controller *TodoController) AddTodo(writer http.ResponseWriter, request *http.Request) {
	var newTodo models.Todo

	if err := json.NewDecoder(request.Body).Decode(&newTodo); err != nil {
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}

	createdTodo, err := controller.Service.AddTodo(newTodo)
	if err != nil {
		http.Error(writer, "Bad request", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(createdTodo)
}

// MARK: - Handle PATCH
func (controller *TodoController) ToggleTodoStatus(writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.Error(writer, "Missing ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(writer, "Ogilitigt ID", http.StatusBadRequest)
	}

	updatedTodo, err := controller.Service.ToggleTodoStatus(id)

	if err != nil {
		http.Error(writer, "Todo not found", http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedTodo)
}
