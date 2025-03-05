package handler

import (
	"example/latihan-golang-gin/internal/repository"
	"example/latihan-golang-gin/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	repo := repository.NewUserRepo()
	usecase := usecase.NewUserUseCase(repo)

	users, err := usecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", GetUsers)
}
