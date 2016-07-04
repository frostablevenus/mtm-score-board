package routes

import (
	"github.com/gin-gonic/gin"

	"mtmScoreBoard/core/config"
	"mtmScoreBoard/core/handlers"
	"mtmScoreBoard/core/middleware"
	"mtmScoreBoard/resources"
)

var (
	playthroughHandler handlers.PlaythroughHandler
)

type Router struct {
}

func CreateEngine(r *resources.Resource) *gin.Engine {
	// Setup resource
	playthroughHandler = handlers.NewPlaythroughHandler(r)

	// Set up gin
	gin.SetMode(config.AppMode)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(middleware.CORS())

	// Setup router
	router := &Router{}
	ApplyRoutes(engine, router)
	return engine
}
