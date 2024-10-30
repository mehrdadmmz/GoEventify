package middlewares

import (
	"REST_PROJECT/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Authenticate is a middleware function that validates the token in the request header.
// If the token is valid, it sets the user ID in the context and calls the next handler.
// If the token is not valid or not present, it aborts the request and sends an unauthorized status.
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
