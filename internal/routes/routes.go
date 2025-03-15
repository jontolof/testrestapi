package routes

import (
	"net/http"
	"testrestapi/internal/controllers"
	"testrestapi/internal/db"
	"testrestapi/internal/repository"
	"testrestapi/internal/services"
)

func SetupRoutes(mux *http.ServeMux) {

	// Initiate layer:repository, service and controller
	todoRepository := repository.NewTodoRepository(db.DB)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	mux.HandleFunc("/todos", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			todoController.GetTodos(writer, request)
		case http.MethodPost:
			todoController.AddTodo(writer, request)
		default:
			http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// For endpoints with an ID (e.g., /todos/{id})
	mux.HandleFunc("/todos/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			todoController.GetTodo(writer, request)
		case http.MethodPatch:
			todoController.ToggleTodoStatus(writer, request)
		default:
			http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
