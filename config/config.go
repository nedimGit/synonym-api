package config

import (
	cfModels "github.com/NedimUka/synonyms/models/config"
)

// AppConfig contains application configuration
var AppConfig cfModels.Config

// Load application configuration
func Load() {

	AppConfig = cfModels.Config{

		Service: cfModels.Service{
			Port:        getEnv("SYNONYMS_SERVICE_PORT", "5500"),
			HTTPTimeout: (getEnvInt("SYNONYMS_HTTP_TIMEOUT", 10)),
		},
	}
}
