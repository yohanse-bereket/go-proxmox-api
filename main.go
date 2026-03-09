package main

import (
	"proxmox-api/config"
	"proxmox-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	router := gin.Default()

	router.POST("/containers", handlers.CreateContainer)
	router.DELETE("/containers/:uuid", handlers.DeleteContainer)
	router.GET("/containers", handlers.ListContainers)

	router.Run(":8080")
}