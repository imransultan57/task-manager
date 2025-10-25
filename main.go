package main

import (
	"log"
	"os"

	"task-manager/config"
	"task-manager/controllers"
	"task-manager/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	// Initialize task collection here
	controllers.InitTaskCollection(config.DB)

	router := gin.Default()
	routes.TaskRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	log.Println("Server running on port", port)
	router.Run(":" + port)
}
