package dbs

import (
	"context"
	"log"
	"time"
	"volga-core/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client = ConnectMongoDB()

func ConnectMongoDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.GetMongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetCollection(name string, indexes bson.D) *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := MongoDB.Database(configs.GetMongoDatabase()).Collection(name)

	mod := mongo.IndexModel{
		Keys:    indexes,
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, mod)

	if err != nil {
		log.Fatal(err)
	}

	return collection
}
