package server

import (
	"github.com/aodai/heimdall/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter creates a new router and returns a pointer.
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	v1 := router.Group("/v1")
	{
		rconsole := new(controllers.RConsoleController)
		v1.GET("/", rconsole.Stats)
		v1.POST("/", rconsole.Vote)
	}
	return router
}
