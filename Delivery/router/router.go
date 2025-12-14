package routers

import (
	"taskmanager/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tc *controllers.TaskController) *gin.Engine {
	r := gin.Default()
	r.POST("/tasks", tc.Create)
	return r
}
