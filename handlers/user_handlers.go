package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/models"
)

func SignUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payloaf"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user created successfully"})

}
