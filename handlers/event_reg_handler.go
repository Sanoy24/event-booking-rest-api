package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sanoy24/event-booking-rest-api/models"
)

func RegisterForEvent(ctx *gin.Context) {
	uerId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event not found"})
		return
	}

	err = event.Register(uerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event not found"})
		return
	}

}

func CancelEventRegistration(ctx *gin.Context) {

}
