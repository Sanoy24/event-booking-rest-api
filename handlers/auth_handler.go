package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/models"
	"github.com/sanoy24/event-booking-rest-api/utils"
)

func Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "incorrect username or password"})
		return
	}

	fmt.Println(user)

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not authenticate user"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "loggedin succesfully", "token": token})
}
