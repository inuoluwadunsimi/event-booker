package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inuoluwadunsimi/event-booker/models"
	"net/http"
)

func main() {
	server := gin.Default()

	server.POST("/events", createEvent)
	server.GET("/events", getEvents)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()

	ctx.JSON(http.StatusOK, gin.H{"message": "yeah?", "data": events})

}

func createEvent(ctx *gin.Context) {

	var event models.Event

	err := ctx.ShouldBind(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse JSON"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}
