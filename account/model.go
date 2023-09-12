package account

import "time"

type Account struct {
	Name        string    `json:"name" bson:"name"`
	Username    string    `json:"username" bson:"username"`
	Password    string    `json:"password,omitempty" bson:"password"`
	Application string    `json:"application,omitempty" bson:"application"`
	User        string    `json:"-" bson:"user"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
