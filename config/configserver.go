package config

// TravelgateX Microservices
// jwt-tools-online - Configuration environment
// Server config private methods

import (
	"github.com/spf13/viper"
)

// Server configuration
type serverConfig struct {
	dev  bool
	host string
	port string
}

// fetch server config
func fetchServerConfig() serverConfig {
	return serverConfig{
		dev:  viper.GetBool("server.dev"),
		host: viper.GetString("server.host"),
		port: viper.GetString("server.port"),
	}
}
