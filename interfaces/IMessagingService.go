package interfaces

import (
	"github.developer.allianz.io/hexalite/fe-messaging-server/models"
)

type IMessagingService interface {
	CreateMessage(payload models.MessagingModel) (bool, error)
	FindMessage(guid string) (*models.MessagingModel, error)
}
