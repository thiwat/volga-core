package user

import (
	"net/http"
	"volga-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func RestRouteV1(router fiber.Router) {

	router.Post("/auth/login", func(c *fiber.Ctx) error {
		var input AuthenInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := AuthLogin(input.Username)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "auth_error",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	router.Post("/auth/validate", func(c *fiber.Ctx) error {
		var input ValidateInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := ValidateOtp(input.Username, input.OTP)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "validate_error",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	router.Patch("/token/validate", func(c *fiber.Ctx) error {
		var input ValidateTokenInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res := ValidateToken(input.Token)

		return c.Status(http.StatusOK).JSON(res)
	})
}
