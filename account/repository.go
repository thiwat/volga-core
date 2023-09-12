package account

import (
	"context"
	"time"
	"volga-core/configs"
	"volga-core/dbs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountCollection *mongo.Collection = dbs.GetCollection(
	configs.GetMongoAccountCollectionName(),
	bson.D{
		{Key: "username", Value: 1},
		{Key: "application", Value: 1},
		{Key: "user", Value: 1},
	},
)

func FindOne(filter bson.M) (Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var account Account

	err := accountCollection.FindOne(ctx, filter).Decode(&account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func Create(account Account) (Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	account.CreatedAt = time.Now()

	_, err := accountCollection.InsertOne(ctx, account)

	if err != nil {
		return account, err
	}

	return FindOne(bson.M{
		"username":    account.Username,
		"application": account.Application,
		"user":        account.User,
	})
}
