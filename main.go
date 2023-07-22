package main

import (
	"JWT-Authentication/database"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	database.MigrateDatabase()

	r.Run()
}
