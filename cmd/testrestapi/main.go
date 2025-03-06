package main

// MARK: - Imports
import (
	"testrestapi/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Ladda in våra routes
	routes.SetupRoutes(router)

	// Starta servern på port 8080
	router.Run("0.0.0.0:9090")
}
