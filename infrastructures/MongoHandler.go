package infrastructures

import (
	log "github.com/sirupsen/logrus"
	"github.developer.allianz.io/hexalite/fe-messaging-server/config"
	"github.developer.allianz.io/hexalite/fe-messaging-server/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoHandler struct {
	Configuration *config.Configuration
}

func (m *MongoHandler) Create(message models.MessagingModel) (bool, error) {
	session, err := mgo.Dial(m.Configuration.Mongo.Server)
	if err != nil {
		log.Info("Failed to establish connection to Mongo server:", err)
		return false, err
	}
	// closes the connection before leaving the function
	defer session.Close()

	// perform operations
	err = session.DB(m.Configuration.Mongo.Database).C(m.Configuration.Mongo.Collection[0]).Insert(message)
	if err != nil {
		log.Info("Failed to insert data:", err)
		return false, err
	}

	return true, nil
}

func (m *MongoHandler) FindByID(guid string) (*models.MessagingModel, error) {

	var messaging *models.MessagingModel = &models.MessagingModel{}

	session, err := mgo.Dial(m.Configuration.Mongo.Server)
	if err != nil {
		log.Info("Failed to establish connection to Mongo server:", err)
		return nil, err
	}
	// closes the connection before leaving the function
	defer session.Close()

	// perform operations
	query := bson.M{
		"$text": bson.M{
			"$search": guid,
		},
	}

	err = session.DB(m.Configuration.Mongo.Database).C(m.Configuration.Mongo.Collection[0]).Find(query).One(&messaging)
	if err != nil {
		log.Info("Failed to fetch the message:", guid, err)
		return nil, err
	}

	return messaging, err
}
