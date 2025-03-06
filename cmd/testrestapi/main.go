package main

// MARK: - Imports
import (
	"testrestapi/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Setup Routes
	routes.SetupRoutes(router)

	// Start servern on port 8080
	router.Run("0.0.0.0:9090")
}
