package account

type Account struct {
	Name      string `json:"name" bson:"name"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"-" bson:"password"`
	User      string `json:"-" bson:"user"`
	CreatedAt string `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at" bson:"updated_at,omitempty"`
}
