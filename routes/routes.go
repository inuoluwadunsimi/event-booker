package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.POST("/events", createEvent)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.PUT("events/:id", updateEvent)

}
