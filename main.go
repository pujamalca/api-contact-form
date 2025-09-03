// Package main serves as the entry point for the API Contact Form application.
//
// It initializes the necessary configurations, sets up the database connection,
// configures repositories, services, and handlers, and starts the HTTP server
// using the Gin framework.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package main

import (
	"api-contact-form/config"
	"api-contact-form/handlers"
	"api-contact-form/helpers"
	"api-contact-form/repositories"
	"api-contact-form/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main is the entry point of the application.
// It performs the following steps:
// 1. Loads environment variables from the .env file.
// 2. Initializes the database connection.
// 3. Sets up repositories, services, and handlers.
// 4. Configures the Gin router with necessary middleware and routes.
// 5. Starts the HTTP server on the specified port.
func main() {
	// Load environment variables from the .env file.
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize the database connection.
	config.InitDB()

	// Initialize repositories, services, and handlers.
	mainHandler := handlers.NewMainHandler()
	healthHandler := handlers.NewHealthHandler()
	contactRepository := repositories.NewContactRepository(config.DB)
	contactService := services.NewContactService(contactRepository)
	contactHandler := handlers.NewContactHandler(contactService)

	// Create a new Gin router with default middleware (logger and recovery).
	router := gin.Default()

	// Configure CORS (Cross-Origin Resource Sharing) settings.
	corsConfig := cors.Config{
		AllowOrigins:     helpers.ParseEnvList("CORS_ALLOWED_ORIGINS"),
		AllowMethods:     helpers.ParseEnvList("CORS_ALLOWED_METHODS"),
		AllowHeaders:     helpers.ParseEnvList("CORS_ALLOWED_HEADERS"),
		AllowCredentials: helpers.GetEnvBool("CORS_ALLOW_CREDENTIALS", false),
		ExposeHeaders:    helpers.ParseEnvList("CORS_EXPOSE_HEADERS"),
		MaxAge:           12 * 60 * 60, // 12 hours
	}

	// Apply the CORS middleware to the router.
	router.Use(cors.New(corsConfig))

	// Define application routes and associate them with their respective handlers.
	router.GET("/", mainHandler.MainHandler)
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/contacts", contactHandler.GetContacts)
	router.GET("/contacts/:id", contactHandler.GetContact)
	router.POST("/contacts", contactHandler.CreateContact)
	router.PUT("/contacts/:id", contactHandler.UpdateContact)
	router.DELETE("/contacts/:id", contactHandler.DeleteContact)

	// Retrieve the application port from environment variables with a default value of "8080".
	appPort := config.GetEnv("APP_PORT", "8080")

	// Start the HTTP server on the specified port.
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}