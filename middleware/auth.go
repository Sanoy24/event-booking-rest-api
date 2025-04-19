package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	fmt.Println(token)
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not Authorized"})
		ctx.Abort()
		return
	}

	parsedToken := strings.SplitN(token, " ", 2)
	fmt.Println(parsedToken)

	if len(parsedToken) != 2 || parsedToken[0] != "Bearer" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid header format"})
		ctx.Abort()
		return
	}
	token = parsedToken[1]

	userId, err := utils.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not Authorized"})
		ctx.Abort()
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()
}
