package user

import (
	"encoding/json"
	"errors"
	"volga-core/dbs"
	"volga-core/types"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
)

var Issuer = "Volga - Password Management"
var SESSION_TTL = 15

func AuthLogin(username string) (AuthenResult, error) {
	_, err := FindByUsername(username)

	if err == nil {
		return AuthenResult{Success: true}, nil
	}

	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      Issuer,
		AccountName: username,
	})

	user := User{
		Username:  username,
		SecretKey: key.Secret(),
	}

	user, err = Create(user)

	if err != nil {
		return AuthenResult{Success: false}, err
	}

	return AuthenResult{
		QRCode: "otpauth://totp/" + Issuer + ":" + username + "?secret=" + key.Secret() + "&issuer=" + Issuer,
	}, nil
}

func ValidateOtp(username string, otp string) (TokenResult, error) {
	var result TokenResult

	user, err := FindByUsername(username)

	if err != nil {
		return result, nil
	}

	valid := totp.Validate(otp, user.SecretKey)

	if !valid {
		return result, errors.New("invalid_otp")
	}

	token := uuid.New().String()
	result.Token = token
	session := types.Session{Username: username, SecretKey: user.SecretKey}

	out, _ := json.Marshal(session)

	dbs.SetKey(token, string(out), SESSION_TTL)

	return result, nil
}

func ValidateToken(token string) ValidateTokenOutput {
	result := ValidateTokenOutput{
		Status: true,
	}

	var session types.Session
	redisData, _ := dbs.GetKey(token)

	if err := json.Unmarshal([]byte(redisData), &session); err != nil {
		result.Status = false
	}

	return result
}
