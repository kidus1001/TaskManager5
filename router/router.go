package router

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", func(ctx *gin.Context) {
		controllers.GetAllTasks(ctx)
	})

	router.GET("/tasks/:id", func(ctx *gin.Context) {
		controllers.GetSpecificData(ctx)
	})

	router.PUT("/tasks/:id", func(ctx *gin.Context) {
		controllers.UpdateSpecificData(ctx)
	})

	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
		controllers.Delete(ctx)
	})

	router.POST("/tasks", func(ctx *gin.Context) {
		controllers.Post(ctx)
	})

	return router
}
