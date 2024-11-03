package routs

import (
	"REST_PROJECT/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// registerForEvent registers a user for a specific event
func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found"})
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Registration failed"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

// cancelRegistration cancels a user's registration for a specific event
func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cancellation failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled"})
}
