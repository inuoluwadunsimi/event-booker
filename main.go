package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inuoluwadunsimi/event-booker/db"
	"github.com/inuoluwadunsimi/event-booker/models"
	"net/http"
	"strconv"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.POST("/events", createEvent)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
	}

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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create events"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}

func getEvent(ctx *gin.Context) {

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid event Id"})

	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
	}

	ctx.JSON(http.StatusOK, event)

}
