package controller

import (
	"net/http"
	"strconv"

	"kanban-kita-api/models"
	"kanban-kita-api/repository"

	"github.com/gin-gonic/gin"
)

type InputTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	CategoryId  uint   `json:"category_id"`
}

func (db Handlers) GetAllTask(c *gin.Context) {
	userData := c.MustGet("id")
	userId := uint(userData.(float64))
	res, err := repository.GetAllTask(db.Connect, userId)
	var result gin.H
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"data": res,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) CreateTask(c *gin.Context) {
	var (
		task   models.Task
		result gin.H
	)
	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	userData := c.MustGet("id")
	userId := uint(userData.(float64))
	task.UserId = userId

	err := repository.CreateTask(task, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"id":          task.ID,
		"title":       task.Title,
		"status":      task.Status,
		"description": task.Description,
		"user_id":     task.UserId,
		"category_id": task.CategoryId,
		"created_at":  task.CreatedAt,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) UpdateTask(c *gin.Context) {
	var (
		reqTask InputTask
		result  gin.H
	)
	if err := c.ShouldBindJSON(&reqTask); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	task := models.Task{
		Title:       reqTask.Title,
		Description: reqTask.Description,
	}
	//task.Description = reqTask.Description

	taskId, _ := strconv.Atoi(c.Param("id"))
	userData := c.MustGet("id")
	userId := uint(userData.(float64))
	_, err := repository.UpdateTask(taskId, userId, task, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"data": task,
	}
	c.JSON(http.StatusCreated, result)
}

func (db Handlers) DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	id, _ := strconv.Atoi(taskId)
	err := repository.DeleteTask(id, db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task has been succesfully deleted",
	})
}

func (db Handlers) UpdateStatusTask(c *gin.Context) {
	var (
		reqTask InputTask
		result  gin.H
	)
	if err := c.ShouldBindJSON(&reqTask); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	task := models.Task{
		Status: reqTask.Status,
	}
	//task.Description = reqTask.Description

	taskId, _ := strconv.Atoi(c.Param("id"))
	userData := c.MustGet("id")
	userId := uint(userData.(float64))
	_, err := repository.UpdateTask(taskId, userId, task, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"data": task,
	}
	c.JSON(http.StatusCreated, result)
}

func (db Handlers) UpdateCategoryTask(c *gin.Context) {
	var (
		reqTask InputTask
		result  gin.H
	)
	if err := c.ShouldBindJSON(&reqTask); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	taskId, _ := strconv.Atoi(c.Param("id"))
	categoryId := reqTask.CategoryId
	category, err := repository.GetCategoryById(db.Connect, int(categoryId))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err = repository.UpdateCategoryTask(db.Connect, taskId, category.ID)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"message": "Update success",
	}
	c.JSON(http.StatusCreated, result)
}
