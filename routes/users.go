package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inuoluwadunsimi/event-booker/models"
	"net/http"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBind(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}
