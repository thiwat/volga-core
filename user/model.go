package user

import "time"

type User struct {
	Username  string    `json:"username" bson:"username,omitempty"`
	SecretKey string    `json:"secret_key" bson:"secret_key,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
