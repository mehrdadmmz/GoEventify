package routes

import (
	"REST_PROJECT/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRouts registers all routes for the application
// it receives a pointer to a gin.Engine instance
func RegisterRouts(server *gin.Engine) {

	// Event routes
	server.GET("/events", getEvents)    // Get all events
	server.GET("/events/:id", getEvent) // Get a specific event by ID

	// Group of authenticated routes
	authenticated := server.Group("/")               // "/" is the base path for all routes in this group
	authenticated.Use(middlewares.Authenticate)      // Use authentication middleware
	authenticated.POST("/events", createEvent)       // Create a new event
	authenticated.PUT("/events/:id", updateEvent)    // Update a specific event by ID
	authenticated.DELETE("/events/:id", deleteEvent) // Delete a specific event by ID

	// Register routes
	authenticated.POST("/events/:id/register", registerForEvent)     // Register for a specific event by ID
	authenticated.DELETE("/events/:id/register", cancelRegistration) // Cancel registration for a specific event by ID

	// User routes
	server.POST("signup", signup) // User signup
	server.POST("/login", login)  // User login
}
