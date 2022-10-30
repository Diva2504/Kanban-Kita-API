package controller

import (
	"fmt"
	"net/http"

	"kanban-kita-api/middleware"
	"kanban-kita-api/models"
	"kanban-kita-api/repository"

	"github.com/gin-gonic/gin"
)

func (db Handlers) UserRegister(c *gin.Context) {
	var data models.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	res, err := repository.CreateUser(&data, db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Cannot register user : %s", err),
		})
	}
	token, err := middleware.GenerateToken(res.ID, res.Email)
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration Succeed",
		"data":    res,
		"token":   token,
	})
}

func (db Handlers) UserLogin(c *gin.Context) {
	var data models.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := repository.UserLogin(&data, db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Cannot Log in User : %s", err),
		})
		return
	}
	token, err := middleware.GenerateToken(res.ID, res.Email)
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Succeed",
		"token":   token,
	})
	return
}

func (db Handlers) DeleteUser(c *gin.Context) {

	userData := c.MustGet("id")
	userId := uint(userData.(float64))

	err := repository.DeleteUser(db.Connect, userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "account deleted",
	})
}

func (db Handlers) UpdateUser(c *gin.Context) {
	var data models.User

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userData := c.MustGet("id")
	userId := uint(userData.(float64))

	err := repository.UpdateUser(db.Connect, int(userId), data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "User Updated",
		"User":   data.FullName,
	})
}
