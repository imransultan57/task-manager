package routes

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	taskRoutes := router.Group("/tasks")

	// ✅ Protect all task routes with JWT middleware
	taskRoutes.Use(middleware.AuthMiddleware())

	{
		taskRoutes.POST("", controllers.CreateTask)
		taskRoutes.GET("", controllers.GetTasks)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
}
