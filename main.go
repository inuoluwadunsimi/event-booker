package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inuoluwadunsimi/event-booker/db"
	"github.com/inuoluwadunsimi/event-booker/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
