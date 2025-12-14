package router

import (
	"taskmanager/controllers"
	"taskmanager/middleware"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", controllers.GetAllTasks)
		auth.GET("/tasks/:id", controllers.GetSpecificData)
	}
	admin := router.Group("/")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.POST("/tasks", controllers.Post)
		admin.PUT("/tasks/:id", controllers.UpdateSpecificData)
		admin.DELETE("/tasks/:id", controllers.Delete)

		admin.POST("/promote", controllers.Promote)
	}

	return router

}
