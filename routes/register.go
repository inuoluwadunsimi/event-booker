package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inuoluwadunsimi/event-booker/models"
	"net/http"
	"strconv"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid event Id"})
		return

	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return

	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "successfully registered"})
}

func cancelRegistration(ctx *gin.Context) {

	userId := ctx.GetInt64("userId")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid event Id"})
		return

	}

	var event models.Event

	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel reg  for event"})
		return

	}

	println("ksdjkdjdjkd")

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully cancelled"})

}
