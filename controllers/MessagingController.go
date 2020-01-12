package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.developer.allianz.io/hexalite/fe-messaging-server/interfaces"
	"github.developer.allianz.io/hexalite/fe-messaging-server/models"
	"github.developer.allianz.io/hexalite/fe-messaging-server/viewmodels"
)

// MessagingController object
type MessagingController struct {
	MessagingService interfaces.IMessagingService
	ErrorService     interfaces.IErrorService
}

// Post method will storing the message from the client
// Post is idempotent -> will do it later
func (ctl *MessagingController) Post(w http.ResponseWriter, r *http.Request) error {

	defer r.Body.Close()

	if r.Method != http.MethodPost {
		return ctl.ErrorService.NewError(nil, http.StatusMethodNotAllowed, "Method not allowed.")
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("Request body read error : %v", err)
	}

	var schema models.MessagingModel
	if err = json.Unmarshal(payload, &schema); err != nil {
		return ctl.ErrorService.NewError(err, http.StatusBadRequest, "Bad request : invalid JSON.")
	}

	result, err := ctl.MessagingService.CreateMessage(schema)
	if result == false || err != nil {
		return fmt.Errorf("create message error : %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json, _ := json.MarshalIndent(&viewmodels.HttpResponseVM{
		StatusCode: http.StatusCreated,
		Result:     schema,
	}, "", " ")
	w.Write(json)
	return nil
}
