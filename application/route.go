package application

import (
	"net/http"
	"volga-core/middlewares"
	"volga-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func RestRouteV1(router fiber.Router) {
	application := router.Group("application")
	application.Use(middlewares.ValidateToken())
	application.Post("/", func(c *fiber.Ctx) error {
		session := c.Locals("session").(types.Session)

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

	application.Put("/:code", func(c *fiber.Ctx) error {
		session := c.Locals("session").(types.Session)

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

	application.Get("/", func(c *fiber.Ctx) error {
		session := c.Locals("session").(types.Session)

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
