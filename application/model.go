package application

import "time"

type Application struct {
	Name      string    `json:"name" bson:"name,omitempty"`
	Code      string    `json:"code" bson:"code,omitempty"`
	User      string    `json:"-" bson:"user,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
