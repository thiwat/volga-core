package user

import (
	"encoding/json"
	"errors"
	"volga-core/dbs"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
)

var Issuer = "Volga - Password Management"

func AuthLogin(username string) (AuthenResult, error) {
	_, err := FindUserByUsername(username)

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

	user, err = CreateUser(user)

	if err != nil {
		return AuthenResult{Success: false}, err
	}

	return AuthenResult{
		QRCode: "otpauth://totp/" + Issuer + ":" + username + "?secret=" + key.Secret() + "&issuer=" + Issuer,
	}, nil
}

func ValidateOtp(username string, otp string) (TokenResult, error) {
	var result TokenResult

	user, err := FindUserByUsername(username)

	if err != nil {
		return result, nil
	}

	valid := totp.Validate(otp, user.SecretKey)

	if !valid {
		return result, errors.New("invalid_otp")
	}

	token := uuid.New().String()
	result.Token = token
	session := Session{Username: username}

	out, _ := json.Marshal(session)

	dbs.SetKey(token, string(out))

	return result, nil
}