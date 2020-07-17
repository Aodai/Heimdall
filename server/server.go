package server

import (
	"fmt"

	"github.com/aodai/heimdall/config"
)

func Init() {
	cfg := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf("%s:%d", cfg.GetString("Web.Address"), cfg.GetInt("Web.Port")))
}
