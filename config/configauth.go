package config

// TravelgateX Microservices
// jwt-tools-online - Configuration environment
// Auth0 config private methods

import (
	"github.com/spf13/viper"
)

// Config Imported from go-jwt-tools
type authConfig struct {
	adminGroup       string
	publicKey        string
	ignoreExpiration bool
	dummyToken       string
	groupsClaim      []string
	memberIDClaim    []string
}

// fetch auth config
func fetchAuthConfig() authConfig {
	return authConfig{
		adminGroup:       viper.GetString("auth.admin-group"),
		publicKey:        viper.GetString("auth.public-key"),
		ignoreExpiration: viper.GetBool("auth.ignore-expiration"),
		dummyToken:       viper.GetString("auth.dummy-token"),
		groupsClaim:      viper.GetStringSlice("auth.groups-claim"),
		memberIDClaim:    viper.GetStringSlice("auth.member-id-claim"),
	}
}
