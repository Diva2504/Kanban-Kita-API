package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secret = "mysecret"

// Generate Json Web Token
func VerifyToken(c *gin.Context) (interface{}, error) {
	errCode := errors.New("please use correct authentication")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errCode
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errCode
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errCode
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, errCode
	}

	return token.Claims.(jwt.MapClaims), nil
}
