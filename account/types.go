package account

type PasswordInput struct {
	Username string `json:"username"`
}

type PasswordOutput struct {
	Password string `json:"password"`
}
