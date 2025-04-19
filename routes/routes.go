package routes

import "github.com/gin-gonic/gin"

func InitRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)
	server.POST("/events", createEvent)
	server.PATCH("/events/:id", updateEvent)
}
