package routs

import (
	"REST_PROJECT/models"
	"REST_PROJECT/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// signup handles user registration
func signup(context *gin.Context) {
	var user models.User

	// Parse request data into user struct instance using ShouldBindJSON method which binds the request body to the user struct
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

// login handles user authentication
func login(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// Validate user credentials using the ValidateCredentials method of the user struct
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	// Generate a token for the user using the GenerateToken function from the utils package so that the user can be authenticated for future requests
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
