package main

import (
	"log"
	"os"

	"task-manager/config"
	"task-manager/controllers"
	"task-manager/routes"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	// Initialize task collection here
	controllers.InitTaskCollection(config.DB)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve static files (CSS, JS)
	router.Static("/static", "./static")

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	routes.TaskRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	log.Println("Server running on port", port)
	router.Run(":" + port)
}
