package middleware

import (
	"net/http"
	"time"

	"kanban-kita-api/config"
	"kanban-kita-api/models"
	"kanban-kita-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secret = "mysecret"

func GenerateToken(id uint, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifiedToken, err := utils.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err":     "unauthorized",
				"message": err.Error(),
			})
		}
		data := verifiedToken.(jwt.MapClaims)
		c.Set("id", data["id"])
		c.Set("email", data["email"])
		c.Set("user_data", verifiedToken)
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.User
		db := config.GetDB()
		userData := c.MustGet("user_data").(jwt.MapClaims)
		userId := int(userData["id"].(float64))
		err := db.Select("role").First(&data, int(userId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data not exist",
			})
			return
		}
		if data.Role != "Admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "don't have access",
			})
			return
		}
		c.Next()
	}
}
