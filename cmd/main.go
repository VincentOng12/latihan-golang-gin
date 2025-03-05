package main

import (
	"example/latihan-golang-gin/config"
	"example/latihan-golang-gin/models"
	"example/latihan-golang-gin/routes"
	"log"
)

func main() {
	config.ConnectDatabase()

	// Auto Migrate Models
	config.DB.AutoMigrate(&models.User{})

	router := routes.SetupRouter()

	port := config.GetEnv("PORT", "7777")
	log.Printf("Server running on :%s", port)
	router.Run(":" + port)
}
