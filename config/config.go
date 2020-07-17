package config

import (
	"github.com/spf13/viper"
)

var cfg *viper.Viper

func init() {
	cfg = viper.New()
	cfg.SetConfigName("conf")
	cfg.SetConfigType("json")
	cfg.AddConfigPath(".")
	err := cfg.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// GetConfig returns a pointer to this object
func GetConfig() *viper.Viper {
	return cfg
}
