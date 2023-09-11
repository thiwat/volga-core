package user

import (
	"context"
	"time"
	"volga-core/configs"
	"volga-core/dbs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = dbs.GetCollection(
	configs.GetMongoUserCollectionName(),
	bson.D{
		{Key: "username", Value: 1},
	},
)

func FindByUsername(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user User

	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func Create(user User) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.CreatedAt = time.Now()
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}

	user, _ = FindByUsername(user.Username)
	return user, nil
}
