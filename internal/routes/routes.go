package routes

import (
	"testrestapi/internal/controllers"
	"testrestapi/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Initiate our service
	todoService := services.NewTodoService()

	// Initiate the controller with the service
	todoController := controllers.NewTodoController(todoService)

	todos := router.Group("/todos")
	{
		todos.GET("", todoController.GetTodos)
		todos.POST("", todoController.AddTodo)
		todos.PATCH("/:id", todoController.ToggleTodoStatus)
	}

	router.GET("/todo/:id", todoController.GetTodo)
}
