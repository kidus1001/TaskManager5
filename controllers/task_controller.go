package controllers

import (
	"net/http"
	"taskmanager/data"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(ctx *gin.Context) {
	all, err := data.GetData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": all})
}

func GetSpecificData(ctx *gin.Context) {
	id := ctx.Param("id")
	task, found := data.GetSpecificData(id)

	if found != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": task})

}

func UpdateSpecificData(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask models.Task
	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := data.UpdateSpecificData(id, updatedTask); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	flag := data.Delete(id)

	if flag != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}

func Post(ctx *gin.Context) {

	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newTask.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "title is required"})
		return
	}

	if err := data.Post(newTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}
