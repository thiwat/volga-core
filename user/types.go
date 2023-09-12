package user

type AuthenInput struct {
	Username string `json:"username"`
}

type AuthenResult struct {
	Success bool   `json:"success"`
	QRCode  string `json:"qrcode_url"`
}

type ValidateInput struct {
	Username string `json:"username"`
	OTP      string `json:"otp"`
}

type TokenResult struct {
	Token string `json:"token"`
}

type ValidateTokenInput struct {
	Token string `json:"token"`
}

type ValidateTokenOutput struct {
	Status bool `json:"status"`
}
