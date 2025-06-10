// routes/routes.go
package routes

import (
	"url-shortener/internal/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine      *gin.Engine
	controllers *controllers.Controllers
}

func NewRouter(engine *gin.Engine, controllers *controllers.Controllers) *Router {
	return &Router{
		engine:      engine,
		controllers: controllers,
	}
}

func (r *Router) RegisterRoutes() {
	// URL routes
	r.engine.POST("/shorten", r.controllers.URL.Shorten)
	r.engine.GET("/:slug", r.controllers.URL.Redirect)

	// User routes
	// r.engine.POST("/register", r.controllers.User.Register)
	// r.engine.POST("/login", r.controllers.User.Login)
}
