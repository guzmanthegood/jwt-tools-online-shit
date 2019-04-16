package config

// TravelgateX Microservices
// jwt-tools-online - Configuration environment
//
// [DEPLOY_MODE]
//  - localDev:  local execution, dev environment
//  - localProd: local execution, prod environment
//  - dev:       cloud execution, dev environment
//  - prod:      cloud execution, prod environment

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/travelgateX/go-jwt-tools/jwt"
)

// AppConfig app config struct
type AppConfig struct {
	server serverConfig // Server basics
	auth   authConfig   // Auth0
}

// LoadConfiguration load configuration file scoped by environmnet
func LoadConfiguration(env string) (*AppConfig, error) {
	log.Printf("[INFO] loading configuration file [./config/env/config_%v] ...\n", env)

	if env == "" {
		return nil, fmt.Errorf("[ERRO] no deploy mode variable given")
	}
	fileName := "config_" + env

	// Viper init
	viper.SetConfigName(fileName)
	viper.AddConfigPath("../config/env") // config default path
	viper.AddConfigPath("./config/env")  // config default path

	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("[ERRO] load configuration error. %s", err)
	}

	log.Println("[INFO] configuration file loaded OK")
	return fetchCurrentConfig(), nil
}

// Fetch current config
func fetchCurrentConfig() *AppConfig {
	return &AppConfig{
		server: fetchServerConfig(),
		auth:   fetchAuthConfig(),
	}
}

// GetJwtParserConfig returns jwt-tools parser config struct
func (a *AppConfig) GetJwtParserConfig() jwt.ParserConfig {
	return jwt.ParserConfig{
		AdminGroup:       a.auth.adminGroup,
		PublicKey:        a.auth.publicKey,
		DummyToken:       a.auth.dummyToken,
		IgnoreExpiration: a.auth.ignoreExpiration,
		GroupsClaim:      a.auth.groupsClaim,
		MemberIDClaim:    a.auth.memberIDClaim,
	}
}

// GetServerHost returns server host
func (a *AppConfig) GetServerHost() string {
	return a.server.host
}

// GetServerPort returns server host
func (a *AppConfig) GetServerPort() string {
	return a.server.port
}

// GetServerAddress returns server address (helper)
func (a *AppConfig) GetServerAddress() string {
	return fmt.Sprintf("http://%s:%s/", a.GetServerHost(), a.GetServerPort())
}

// GetAPIEndPoint returns server address (helper)
func (a *AppConfig) GetAPIEndPoint() string {
	return fmt.Sprintf("%sapi/v1/", a.GetServerAddress())
}

// IsDevelopment returns true if development environment
func (a *AppConfig) IsDevelopment() bool {
	return a.server.dev
}
