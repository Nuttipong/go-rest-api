package interfaces

import "github.developer.allianz.io/hexalite/fe-messaging-server/models"

type IMongoRepository interface {
	Create(message models.MessagingModel) (bool, error)
	FindByID(guid string) (*models.MessagingModel, error)
}
