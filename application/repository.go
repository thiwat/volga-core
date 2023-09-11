package application

import (
	"context"
	"time"
	"volga-core/configs"
	"volga-core/dbs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var applicationCollection *mongo.Collection = dbs.GetCollection(
	configs.GetMongoApplicationCollectionName(),
	bson.D{
		{Key: "code", Value: 1},
		{Key: "user", Value: 1},
	},
)

func FindByCode(code string, user string) (Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var app Application

	err := applicationCollection.FindOne(
		ctx,
		bson.M{"user": user, "code": code},
	).Decode(&app)

	if err != nil {
		return app, err
	}

	return app, nil
}

func Create(app Application) (Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.CreatedAt = time.Now()
	_, err := applicationCollection.InsertOne(ctx, app)
	if err != nil {
		return app, err
	}
	return FindByCode(app.Code, app.User)
}

func UpdateByCode(code string, app Application) (Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.UpdatedAt = time.Now()

	_, err := applicationCollection.UpdateOne(
		ctx,
		bson.M{"code": code, "user": app.User},
		bson.M{"$set": app},
	)

	if err != nil {
		return app, err
	}
	return FindByCode(code, app.User)
}

func ListByUser(user string) ([]Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var apps = make([]Application, 0)

	opts := options.Find()

	cursor, err := applicationCollection.Find(
		ctx,
		bson.M{"user": user},
		opts,
	)

	if err != nil {
		return apps, err
	}

	if err = cursor.All(ctx, &apps); err != nil {
		return apps, err
	}

	return apps, err
}
