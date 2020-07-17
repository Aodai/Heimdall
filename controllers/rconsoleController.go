package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aodai/heimdall/cache"
	"github.com/aodai/heimdall/config"
	"github.com/aodai/heimdall/model"
	"github.com/aodai/heimdall/rconsole"
	"github.com/gin-gonic/gin"
)

// RConsoleController handles functions for RConsole.
type RConsoleController struct{}

var (
	cfg        = config.GetConfig()
	statsModel = new(model.Stats)
	appCache   = cache.GetCache()
)

// Vote handles the vote callback.
func (t RConsoleController) Vote(c *gin.Context) {
	fmt.Printf("IP: %s UID: %s Custom: %s\n", c.PostForm("ip"), c.PostForm("userid"), c.PostForm("custom"))
	c.Status(http.StatusAccepted)
}

// Stats returns general info about a server.
func (t RConsoleController) Stats(c *gin.Context) {
	var stats model.Stats
	var found bool
	data, found := appCache.Get("stats")
	if found {
		stats = data.(model.Stats)
	} else {
		var rc rconsole.RConsole
		rc.Init()
		stats = rc.FetchStats()
		appCache.Set("stats", stats, 30*time.Second)
		rc.Close()
	}
	c.JSON(200, stats)
}
