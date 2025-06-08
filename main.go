package main

import (
	"url-shortener/config"
	"url-shortener/internal/controllers"
	"url-shortener/internal/infrastructure/mongo"
	"url-shortener/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := mongo.NewMongoURLRepository(config.DB)
	usecase := usecases.NewURLUsecase(repo)
	controller := controllers.NewURLController(usecase)

	r := gin.Default()
	r.POST("/shorten", controller.Shorten)
	r.GET("/:slug", controller.Redirect)

	r.Run(":8080")
}
