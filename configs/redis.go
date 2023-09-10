package configs

import (
	"os"
	"strconv"
	"strings"
)

func GetRedisURI() string {
	tmp := strings.Split(os.Getenv("SESSION_REDIS_URI"), "/")
	return strings.Join(tmp[2:len(tmp)-1], "/")
}

func GetRedisDB() int {
	tmp := strings.Split(os.Getenv("SESSION_REDIS_URI"), "/")
	value, _ := strconv.Atoi(tmp[len(tmp)-1])
	return value
}
