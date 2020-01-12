package services

import (
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.developer.allianz.io/hexalite/fe-messaging-server/interfaces"
	"github.developer.allianz.io/hexalite/fe-messaging-server/models"
	"gopkg.in/mgo.v2/bson"
)

type MessagingService struct {
	MongoRepository interfaces.IMongoRepository
}

func (svc *MessagingService) CreateMessage(payload models.MessagingModel) (bool, error) {

	var success bool = false
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	payload.ID = bson.NewObjectId()
	payload.ServerTime = time.Now()

	go func(data models.MessagingModel) {
		defer waitGroup.Done()
		result, err := svc.MongoRepository.Create(data)
		if err != nil {
			log.Warn("Cannot insert message to Mongo\n", err)
		}
		success = result
	}(payload)

	waitGroup.Wait()

	return success, nil
}

func (svc *MessagingService) FindMessage(guid string) (*models.MessagingModel, error) {

	vm, err := svc.MongoRepository.FindByID(guid)
	if err != nil {
		log.Warn(fmt.Sprintf("Cannot found requestID: %s\n", guid), err)
	}

	return vm, err
}
