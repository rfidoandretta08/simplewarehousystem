package main

import (
	"assigment/config"
	"assigment/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	routes.SetupRoutes(router)

	log.Println("Server running on port 8080")
	router.Run(":8080")

}
