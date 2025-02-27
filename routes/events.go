package routs

import (
	"REST_PROJECT/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getEvents fetches all events from the database
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		// If there is an error, return an internal server error which is a 500 status code and a message
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	// If there are no errors, return the events and a 200 status code which means OK
	context.JSON(http.StatusOK, events)
}

// getEvent fetches a single event by its ID
func getEvent(context *gin.Context) {
	// Parse the event ID from the request URL and convert it to an integer using the ParseInt function
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		// If there is an error, return a bad request status code which is 400 and a message
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

// createEvent creates a new event
func createEvent(context *gin.Context) {
	var event models.Event

	// Bind the request data to the event struct using the ShouldBindJSON function which is a method of the context object
	// that binds the request data to the struct and returns an error if the request data is not valid JSON
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// Get the user ID from the context object which is set by the authentication middleware
	userId := context.GetInt64("userId")
	event.UserID = userId // Set the user ID for the event object

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

// updateEvent updates an existing event
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	// Fetch the event from the database, so we can check if the user is authorized to update it
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	// If the user ID in the context object is not the same as the user ID of the event, return a 401 status code which means unauthorized
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update the event."})
		return
	}

	var updatedEvent models.Event
	// bind the request data to the updatedEvent struct using the ShouldBindJSON function
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

// deleteEvent deletes an existing event
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete the event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
