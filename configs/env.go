package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetPort() string {
	return ":" + os.Getenv("PORT")
}
