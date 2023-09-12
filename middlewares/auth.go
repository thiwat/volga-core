package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
	"volga-core/dbs"
	"volga-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {

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

		c.Locals("session", session)

		return c.Next()
	}
}
