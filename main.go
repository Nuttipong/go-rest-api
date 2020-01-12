package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"github.developer.allianz.io/hexalite/fe-messaging-server/config"

	_ "net/http/pprof"
)

var (
	configuration *config.Configuration
	// Version represent version via -ldflag
	Version string = "0.0.0"
	// Time represent build time via -ldflag
	Time string
	// User represent build user via -ldflag
	User string
)

func main() {

	log.Info("Setup config file..")
	configuration = config.NewConfiguration()

	log.Info("Setup routes..")
	router := NewRouter(configuration).InitRouter()

	log.Info("Setup CORS..")
	cors := handlers.CORS(
		handlers.AllowedHeaders(configuration.App.AllowedHeaders),
		handlers.AllowedOrigins(configuration.App.AllowedOrigins),
		handlers.AllowedMethods(configuration.App.AllowedMethods),
	)
	router.Use(cors)

	addr := fmt.Sprintf(":%d", configuration.App.Port)
	log.Info(fmt.Sprintf("Applicatio version: %s", Version))
	log.Info(fmt.Sprintf("Applicatio build time: %s", Time))
	log.Info(fmt.Sprintf("Applicatio build user: %s", User))
	log.Info(fmt.Sprintf("Applicatio started on port %s", addr))
	log.Fatal(http.ListenAndServe(addr, router))
}
