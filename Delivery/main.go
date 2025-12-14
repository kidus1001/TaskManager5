package main

import (
	"taskmanager/Delivery/controllers"
	routers "taskmanager/Delivery/router"
	repositories "taskmanager/Repositories"
	usecases "taskmanager/Usecases"
)

func main() {
	taskRepo := repositories.NewMongoTaskRepository()
	taskUC := usecases.NewTaskUsecase(taskRepo)
	taskCtrl := controllers.NewTaskController(taskUC)

	r := routers.SetupRouter(taskCtrl)
	r.Run(":8080")
}
