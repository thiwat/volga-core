package configs

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func GetMongoURI() string {
	tmp := strings.Split(os.Getenv("MONGODB_URI"), "/")
	return strings.Join(tmp[:len(tmp)-1], "/")
}

func GetMongoDatabase() string {
	tmp := strings.Split(os.Getenv("MONGODB_URI"), "/")
	return tmp[len(tmp)-1]
}

func GetMongoApplicationCollectionName() string {
	return "applications"
}

func GetMongoAccountCollectionName() string {
	return "accounts"
}

func GetMongoUserCollectionName() string {
	return "users"
}
