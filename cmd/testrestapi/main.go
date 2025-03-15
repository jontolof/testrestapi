package main

// MARK: - Imports
import (
	"log"
	"net/http"
	"testrestapi/internal/db"
	"testrestapi/internal/routes"
)

func main() {
	// Initiate database connection (and potential migration)
	db.Init()

	// Create Mux
	mux := http.NewServeMux()

	// Setup Routes
	routes.SetupRoutes(mux)

	// Start servern on port 8080
	log.Println("Server starting on :8080")
	err := http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
