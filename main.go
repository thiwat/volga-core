package main

import (
	"log"
	"volga-core/application"
	"volga-core/configs"
	"volga-core/user"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	user.RestRouteV1(v1)
	application.RestRouteV1(v1)

	log.Fatal(app.Listen(configs.GetPort()))
}
