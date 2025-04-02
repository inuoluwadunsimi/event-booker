package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inuoluwadunsimi/event-booker/models"
	"github.com/inuoluwadunsimi/event-booker/utils"
	"net/http"
	"strconv"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "yeah?", "data": events})

}

func createEvent(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized error"})
		return
	}

	userId, err := utils.VerifToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized error"})
		return
	}

	var event models.Event

	err = ctx.ShouldBind(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse JSON"})
		return
	}

	event.UserID = userId
	err = event.Save()
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

func updateEvent(ctx *gin.Context) {

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid event Id"})

	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBind(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse JSON"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "event updated succesfully"})

}

func deleteEvent(ctx *gin.Context) {

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid event Id"})

	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.Delete()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted succesfully"})
}
