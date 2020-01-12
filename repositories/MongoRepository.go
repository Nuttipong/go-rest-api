package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.developer.allianz.io/hexalite/fe-messaging-server/infrastructures"
	"github.developer.allianz.io/hexalite/fe-messaging-server/interfaces"
	"github.developer.allianz.io/hexalite/fe-messaging-server/models"
)

type MongoRepository struct {
	MongoHandler *infrastructures.MongoHandler
}

func (m *MongoRepository) Create(message models.MessagingModel) (bool, error) {
	return m.MongoHandler.Create(message)
}

func (m *MongoRepository) FindByID(guid string) (*models.MessagingModel, error) {
	return m.MongoHandler.FindByID(guid)
}

type MongoRepositoryWithCircuitBreaker struct {
	MongoRepository interfaces.IMongoRepository
}

func (m *MongoRepositoryWithCircuitBreaker) Create(message models.MessagingModel) (bool, error) {
	output := make(chan bool, 1)
	hystrix.ConfigureCommand("mongo_repo_create_new_message", hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})
	errors := hystrix.Go("mongo_repo_create_new_message", func() error {
		result, err := m.MongoRepository.Create(message)
		if err != nil {
			return err
		}
		output <- result
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return false, err
	}
}

func (m *MongoRepositoryWithCircuitBreaker) FindByID(guid string) (*models.MessagingModel, error) {
	output := make(chan *models.MessagingModel, 1)
	hystrix.ConfigureCommand("mongo_repo_find_message", hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})
	errors := hystrix.Go("mongo_repo_find_message", func() error {
		result, err := m.MongoRepository.FindByID(guid)
		if err != nil {
			return err
		}
		output <- result
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return &models.MessagingModel{}, err
	}
}
