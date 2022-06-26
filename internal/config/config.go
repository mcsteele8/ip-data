package config

import (
	"ip-data/tools/wlog"
	"os"
)

const (
	EnvKeyForDomainApiKey = "DOMAIN_API_KEY"
	EnvKeyForServerPort   = "IP_DATA_SERVER_PORT"
)

var (
	DomainApiKey = ""
	ServerPort   = ":8081"
)

func init() {
	keyValue := os.Getenv(EnvKeyForDomainApiKey)
	if keyValue != "" {
		wlog.New().Infof("Loading %s", EnvKeyForDomainApiKey)
		DomainApiKey = keyValue
	}

	portValue := os.Getenv(EnvKeyForServerPort)
	if portValue != "" {
		wlog.New().Infof("Loading %s", EnvKeyForServerPort)
		ServerPort = portValue
	}

}
