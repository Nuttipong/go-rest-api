package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MessagingModel struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	GUID       string        `json:"guid" bson:"_guid"`
	Data       Data          `json:"data"`
	ServerTime time.Time     `json:"serverTime"`
}

type Data struct {
	ClientType    string    `json:"clientType"`
	MetricName    string    `json:"metricName"`
	Error         string    `json:"error"`
	Origin        string    `json:"origin"`
	Browser       string    `json:"browser"`
	CurrentRoute  string    `json:"currentRoute"`
	PreviousRoute string    `json:"previousRoute"`
	IP            string    `json:"ip"`
	Method        string    `json:"method"`
	OS            string    `json:"os"`
	MobileType    string    `json:"mobileType"`
	ClientTime    time.Time `json:"clientTime"`
}
