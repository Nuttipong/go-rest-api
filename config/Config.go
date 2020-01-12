package config

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration is a wrapper to supply the whole configurations which mean 1 app would be contain 1 config
type Configuration struct {
	App   appConfig
	Mongo mongoConfig
}

type appConfig struct {
	Server         string
	Port           int
	PerfPort       int
	AllowedOrigins []string
	AllowedHeaders []string
	AllowedMethods []string
}

type mongoConfig struct {
	Server     string
	Database   string
	Collection []string
	Password   string
}

var (
	configuration *Configuration
	configOnce    sync.Once
)

// NewConfiguration is a constructor and will create once
func NewConfiguration() *Configuration {

	if configuration == nil {
		configOnce.Do(func() {

			viper.SetConfigName("app")
			viper.SetConfigType("json")
			viper.AddConfigPath(".")
			viper.AutomaticEnv()
			viper.SetDefault("FE_MESSAGING_SERVER_ORIGIN_ALLOWED", "http://example.com")
			viper.SetDefault("FE_MESSAGING_SERVER_ENV", "noneprod")
			viper.SetDefault("FE_MESSAGING_SERVER_MONGO_SERVER", "127.0.0.1:27017")
			viper.SetDefault("FE_MESSAGING_SERVER_MONGO_DATABASE", "messaging_db")
			viper.SetDefault("FE_MESSAGING_SERVER_MONGO_COLLECTION", "client_message")
			viper.SetDefault("FE_MESSAGING_SERVER_MONGO_PASSWORD", "")

			if err := viper.ReadInConfig(); err != nil {
				if err, ok := err.(viper.ConfigFileNotFoundError); ok {
					log.Fatal("App config file not found", err)
				} else {
					log.Fatal("Another error was produced", err)
				}
			}

			appConfig := appConfig{}
			if err := viper.Unmarshal(&appConfig); err != nil {
				log.Fatal("Unable to decode App config struct, %v", err)
			}

			configuration = &Configuration{
				App:   appConfig,
				Mongo: getMongoConfig(),
			}

		})
	}

	return configuration
}

func getMongoConfig() mongoConfig {
	return mongoConfig{
		Server:   viper.GetString("FE_MESSAGING_SERVER_MONGO_SERVER"),
		Database: viper.GetString("FE_MESSAGING_SERVER_MONGO_DATABASE"),
		Collection: []string{
			viper.GetString("FE_MESSAGING_SERVER_MONGO_COLLECTION"),
		},
		Password: viper.GetString("FE_MESSAGING_SERVER_MONGO_PASSWORD"),
	}
}
