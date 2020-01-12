package main

import (
	"sync"

	"github.developer.allianz.io/hexalite/fe-messaging-server/config"
	"github.developer.allianz.io/hexalite/fe-messaging-server/controllers"
	"github.developer.allianz.io/hexalite/fe-messaging-server/infrastructures"
	"github.developer.allianz.io/hexalite/fe-messaging-server/repositories"
	"github.developer.allianz.io/hexalite/fe-messaging-server/services"
)

// ServiceContainer struct
type ServiceContainer struct {
	configuration *config.Configuration
}

// IServiceContainer interfaec
type IServiceContainer interface {
	RegisterMongoHandler() infrastructures.MongoHandler
	RegisterLiveNessController() controllers.LivenessController
	RegisterMessagingController(mongoHandler *infrastructures.MongoHandler) controllers.MessagingController
}

var (
	k             *ServiceContainer
	containerOnce sync.Once
)

// NewServiceContainer is a constructor
func NewServiceContainer(configuration *config.Configuration) IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &ServiceContainer{
				configuration: configuration,
			}
		})
	}
	return k
}

// RegisterMongoHandler to create MongoHandler instance and inject dependency if need
func (k *ServiceContainer) RegisterMongoHandler() infrastructures.MongoHandler {
	return infrastructures.MongoHandler{Configuration: k.configuration}
}

// RegisterLiveNessController to create an instance and inject dependency if need
func (k *ServiceContainer) RegisterLiveNessController() controllers.LivenessController {
	liveNessController := controllers.LivenessController{}
	return liveNessController
}

// RegisterMessagingController to create an instance and inject dependency if need
func (k *ServiceContainer) RegisterMessagingController(mongoHandler *infrastructures.MongoHandler) controllers.MessagingController {

	errorService := services.NewErrorService()
	messageRepository := &repositories.MongoRepository{MongoHandler: mongoHandler}
	messagingService := &services.MessagingService{
		MongoRepository: &repositories.MongoRepositoryWithCircuitBreaker{messageRepository},
	}
	messagingController := controllers.MessagingController{MessagingService: messagingService, ErrorService: errorService}

	return messagingController
}
