package app

import (
	"url-shortener/config"
	"url-shortener/internal/controllers"
	"url-shortener/internal/infrastructure/mongo/repositories"
	"url-shortener/internal/routes"
	"url-shortener/internal/usecases"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	router := gin.Default()

	// Inicializa as dependências
	urlRepo := repositories.NewMongoURLRepository(config.GetMongoDB())
	urlUsecase := usecases.NewURLUsecase(urlRepo)

	// Cria os controllers
	controllers := controllers.NewControllers(urlUsecase)

	// Configura as rotas
	routes.NewRouter(router, controllers).RegisterRoutes()

	return &App{
		Router: router,
	}
}
