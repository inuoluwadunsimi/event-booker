package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inuoluwadunsimi/event-booker/models"
	"github.com/inuoluwadunsimi/event-booker/utils"
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

func login(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request error"})
		return

	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return

	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})

}
