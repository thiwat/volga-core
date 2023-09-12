package account

import (
	"net/http"
	"volga-core/middlewares"
	"volga-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func RestRouteV1(router fiber.Router) {
	account := router.Group("account")

	account.Use(middlewares.ValidateToken())

	account.Post("/", func(c *fiber.Ctx) error {
		session := c.Locals("session").(types.Session)

		var account Account
		if err := c.BodyParser(&account); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := CreateAccount(account, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	account.Get("/:application", func(c *fiber.Ctx) error {
		app := c.Params("application")
		session := c.Locals("session").(types.Session)

		res, err := ListAccount(app, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	account.Patch("/:application", func(c *fiber.Ctx) error {
		session := c.Locals("session").(types.Session)
		app := c.Params("application")

		var input PasswordInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := GetPassword(app, input.Username, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})
}
