package domain

import "time"

type Message struct {
	Data             string    `json:"data" bson:"data"`
	InsertedIn       time.Time `json:"insertedIn" bson:"insertedIn"`
	ApplicationName  string    `json:"applicationName" bson:"applicationName"`
	Squad            string    `json:"squad" bson:"squad"`
	DescriptionError string    `json:"descriptionError" bson:"descriptionError"`
}
