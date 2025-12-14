package controllers

import (
	"net/http"

	domain "taskmanager/Domain"
	usecases "taskmanager/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	usecase *usecases.TaskUsecase
}

func NewTaskController(u *usecases.TaskUsecase) *TaskController {
	return &TaskController{u}
}

func (t *TaskController) Create(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.usecase.Create(task)
	c.JSON(http.StatusCreated, gin.H{"message": "task created"})
}
