package main

import (
	"JWT-Authentication/controllers"
	"JWT-Authentication/database"
	"JWT-Authentication/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.MigrateDatabase()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.Auth, controllers.Validate)
	r.GET("/getnotes", middleware.Auth, controllers.GetNotes)
	r.POST("/insertnote", middleware.Auth,controllers.InsertNote)
	r.GET("/getnote/:title", middleware.Auth,controllers.GetNote)

	r.Run()
}
