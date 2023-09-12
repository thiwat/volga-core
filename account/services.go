package account

import (
	"volga-core/types"
	"volga-core/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAccount(account Account, session types.Session) (Account, error) {
	account.User = session.Username
	account.Password = utils.Encrypt(account.Password, session.SecretKey)

	account, err := Create(account)

	if err != nil {
		return account, err
	}

	account.Password = ""
	account.Application = ""

	return account, nil
}

func ListAccount(app string, session types.Session) ([]Account, error) {
	accounts, err := List(bson.M{
		"user":        session.Username,
		"application": app,
	})

	if err != nil {
		return accounts, err
	}

	for i, _ := range accounts {
		accounts[i].Password = ""
		accounts[i].Application = ""
	}

	return accounts, nil
}

func GetPassword(app string, username string, session types.Session) (PasswordOutput, error) {
	var res PasswordOutput
	account, err := FindOne(bson.M{
		"application": app,
		"username":    username,
		"user":        session.Username,
	})

	if err != nil {
		return res, err
	}

	res.Password = utils.Decrypt(account.Password, session.SecretKey)

	return res, nil
}
