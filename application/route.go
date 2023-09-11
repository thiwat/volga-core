package application

import (
	"encoding/json"
	"net/http"
	"strings"
	"volga-core/dbs"
	"volga-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func RestRouteV1(router fiber.Router) {
	router.Post("/application", func(c *fiber.Ctx) error {
		token := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", 1)

		var session types.Session

		sessionData, _ := dbs.GetKey(token)

		err := json.Unmarshal([]byte(sessionData), &session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: "Access Denided",
			})
		}

		var app Application
		if err := c.BodyParser(&app); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		app.User = session.Username

		res, err := Create(app)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	router.Put("/application/:code", func(c *fiber.Ctx) error {
		token := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", 1)

		var session types.Session

		sessionData, _ := dbs.GetKey(token)

		err := json.Unmarshal([]byte(sessionData), &session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: "Access Denided",
			})
		}

		code := c.Params("code")
		var app Application
		if err := c.BodyParser(&app); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		app.User = session.Username

		res, err := UpdateByCode(code, app)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	router.Get("/application", func(c *fiber.Ctx) error {
		token := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", 1)

		var session types.Session

		sessionData, _ := dbs.GetKey(token)

		err := json.Unmarshal([]byte(sessionData), &session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: "Access Denided",
			})
		}

		res, err := ListByUser(session.Username)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})
}
