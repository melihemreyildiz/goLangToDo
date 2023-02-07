package authentication

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(header, "Bearer ")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("plmsajdBDSAHDHhuen!!.@34??-^^^%aldask34nncBHSDAS"), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("userID", claims["userID"])
		c.Set("email", claims["email"])
		log.Println("[GIN-debug] Listening and serving HTTP on localhost:80802", c)
		c.Next()
	} else {
		log.Println("[GIN-debug] Listening and serving HTTP on localhost:80804")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
}
