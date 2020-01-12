package services

import (
	"encoding/json"
	"fmt"

	"github.developer.allianz.io/hexalite/fe-messaging-server/interfaces"
)

type ErrorService struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *ErrorService) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

func (e *ErrorService) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

func (e *ErrorService) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func NewErrorService() interfaces.IErrorService {
	return &ErrorService{}
}

func (e *ErrorService) NewError(err error, status int, detail string) error {
	return &ErrorService{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}
