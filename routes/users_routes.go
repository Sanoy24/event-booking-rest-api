package routes

import "github.com/gin-gonic/gin"

func usersRoutes(userGroup *gin.RouterGroup) {
	users := userGroup.Group("/users")
	{
		users.POST("/signup", signUp)
	}
}
