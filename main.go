package main

import (
	"golang-test/config"
	"golang-test/middlewares"
	"golang-test/models"
	"golang-test/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Setup validator
	config.SetupValidator()

	// Migrate the schema
	err := config.DB.AutoMigrate(&models.Address{})
	if err != nil {
		log.Fatal("Failed to migrate database schema!", err)
	}

	r := gin.New()

	// Middlewares
	r.Use(gin.Logger(), gin.Recovery(), middlewares.ErrorHandler())

	// Setup routes
	routes.SetupRouter(r)

	// Listen and serve on 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
