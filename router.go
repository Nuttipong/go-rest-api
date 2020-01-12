package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.developer.allianz.io/hexalite/fe-messaging-server/config"
	"github.developer.allianz.io/hexalite/fe-messaging-server/interfaces"
)

type IRouter interface {
	InitRouter() *mux.Router
}

type Router struct {
	configuration *config.Configuration
}

var (
	m          *Router
	routerOnce sync.Once
)

func NewRouter(configuration *config.Configuration) IRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &Router{
				configuration: configuration,
			}
		})
	}
	return m
}

func (router *Router) InitRouter() *mux.Router {

	container := NewServiceContainer(router.configuration)
	mongoHandler := container.RegisterMongoHandler()
	liveNessController := container.RegisterLiveNessController()
	messagingController := container.RegisterMessagingController(&mongoHandler)

	routes := mux.NewRouter().StrictSlash(false)
	routes = routes.PathPrefix("/api/v1/").Subrouter()

	routes.Handle("/message", errorHandler(messagingController.Post))
	routes.HandleFunc("/liveness", liveNessController.Get).Methods("GET")

	return routes
}

// errorHandler is a wrapper aroud handler for any handler.
type errorHandler func(http.ResponseWriter, *http.Request) error

func (fn errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}

	log.Printf("An error accured: %v", err)

	clientError, ok := err.(interfaces.IClientError)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := clientError.ResponseBody()
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status, headers := clientError.ResponseHeaders()
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
	w.Write(body)
}
