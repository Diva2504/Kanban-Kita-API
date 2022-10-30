package controller

import (
	"net/http"
	"strconv"

	"kanban-kita-api/models"
	"kanban-kita-api/repository"

	"github.com/gin-gonic/gin"
)

/*type ResponseCategory struct {
	ID     uint    `json:"id"`
	Type   string `json:"type"`
	Created_at time.Time    `json:"created_at"`
	Updated_at time.Time  `json:"updated-at"`
}*/

/*type InputCategory struct {
	Type string `json:"type"`
}*/

func (db Handlers) GetAllCategory(c *gin.Context) {
	res, err := repository.GetAllCategories(db.Connect)
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

func (db Handlers) CreateCategory(c *gin.Context) {
	var (
		category models.Category
		result   gin.H
	)
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := repository.CreateCategory(category, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"type": category.Type,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) UpdateCategory(c *gin.Context) {
	var (
		category models.Category
		result   gin.H
	)
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	categoryId, _ := strconv.Atoi(c.Param("id"))
	_, err := repository.UpdateCategory(categoryId, category, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"data": category,
	}
	c.JSON(http.StatusCreated, result)
}

func (db Handlers) DeleteCtegory(c *gin.Context) {
	categoryId := c.Param("id")
	id, _ := strconv.Atoi(categoryId)
	err := repository.DeleteCategory(id, db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "category has been succesfully deleted",
	})
}
