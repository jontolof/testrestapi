package routes

import (
	"net/http"
	"testrestapi/internal/controllers"
	"testrestapi/internal/services"
)

func SetupRoutes(mux *http.ServeMux) {
	// Initiate our service
	todoService := services.NewTodoService()

	// Initiate the controller with the service
	todoController := controllers.NewTodoController(todoService)

	mux.HandleFunc("/todos", func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			todoController.GetTodos(responseWriter, request)
		} else if request.Method == http.MethodPost {
			todoController.AddTodo(responseWriter, request)
		} else {
			http.Error(responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// For endpoints with an ID (e.g., /todos/{id})
	mux.HandleFunc("/todos/", func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			todoController.GetTodo(responseWriter, request)
		} else if request.Method == http.MethodPatch {
			todoController.ToggleTodoStatus(responseWriter, request)
		} else {
			http.Error(responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
