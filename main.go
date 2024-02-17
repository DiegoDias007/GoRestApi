package main

import (
	"api.com/controllers"
	"api.com/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	server := gin.Default()
	controllers.RegisterEventController(server)
	server.Run(":8080") // localhost: 8080
}