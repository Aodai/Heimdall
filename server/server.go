package server

import (
	"fmt"

	"github.com/aodai/heimdall/config"
)

// Run starts the web server.
func Run() {
	cfg := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf("%s:%d", cfg.GetString("Web.Address"), cfg.GetInt("Web.Port")))
}
