package routes

import (
	"testrestapi/internal/controllers"
	"testrestapi/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Initiera vår service
	todoService := services.NewTodoService()

	// Initiera controllern med servicen
	todoController := controllers.NewTodoController(todoService)

	todos := router.Group("/todos")
	{
		todos.GET("", todoController.GetTodos)
		todos.POST("", todoController.AddTodo)
		todos.PATCH("/:id", todoController.ToggleTodoStatus)
	}

	router.GET("/todo/:id", todoController.GetTodo)
}
