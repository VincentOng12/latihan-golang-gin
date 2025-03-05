package main

import (
	"example/latihan-golang-gin/api/handler"
	"example/latihan-golang-gin/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load konfigurasi
	config.Load()

	// Setup router Gin
	r := gin.Default()

	// Register routes
	handler.RegisterRoutes(r)

	// Start server
	port := config.GetEnv("PORT", "7777")
	log.Printf("Server running on :%s", port)
	r.Run(":" + port)
}
