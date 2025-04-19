package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/handlers"
)

func usersRoutes(userGroup *gin.RouterGroup) {
	users := userGroup.Group("/users")
	{
		users.POST("/signup", handlers.SignUp)
	}
}
