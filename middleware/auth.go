package middleware

import (
	"JWT-Authentication/database"
	"JWT-Authentication/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Auth(c *gin.Context) {
	// Get the cookie of the request

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//Decode / Validate it
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	// Check the expiration time

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUpgradeRequired)
		}

		//Find the user with token sub
		var user models.User
		DB,_:=database.ConnectToDatabase()
		DB.First(&user,claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("Error---")
		}

		//Attach to request
		c.Set("user", user)

		//Continue;
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
