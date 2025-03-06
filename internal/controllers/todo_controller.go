package controllers

import (
	"net/http"
	"testrestapi/internal/models"
	"testrestapi/internal/services"

	"github.com/gin-gonic/gin"
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
func (controller *TodoController) GetTodos(context *gin.Context) {
	todos, err := controller.Service.GetTodos()
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	context.IndentedJSON(http.StatusOK, todos)
}

func (controller *TodoController) GetTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := controller.Service.GetTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

// MARK: - Handle POST
func (controller *TodoController) AddTodo(context *gin.Context) {
	var newTodo models.Todo

	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	todo, err := controller.Service.AddTodo(&newTodo)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	context.IndentedJSON(http.StatusCreated, todo)
}

// MARK: - Handle PATCH
func (controller *TodoController) ToggleTodoStatus(context *gin.Context) {
	id := context.Param("id")

	todo, err := controller.Service.ToggleTodoStatus(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}
