package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/handlers"
)

func authRoutes(userGroup *gin.RouterGroup) {
	auth := userGroup.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
	}
}
