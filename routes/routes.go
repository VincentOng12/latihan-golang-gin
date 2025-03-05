package routes

import (
	"example/latihan-golang-gin/api/handlers"
	"example/latihan-golang-gin/repositories"
	"example/latihan-golang-gin/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRepo := &repositories.UserRepository{}
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.POST("/users", userHandler.RegisterUser)
	router.GET("/users", userHandler.GetAllUsers)

	return router
}
