package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/inuoluwadunsimi/event-booker/utils"
	"net/http"
)

func Authenticate(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized error"})
		return
	}

	userId, err := utils.VerifToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized error"})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
