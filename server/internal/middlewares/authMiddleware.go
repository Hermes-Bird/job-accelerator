package middlewares

import (
	"github.com/Hermes-Bird/job-accelerator/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
)

func NewAuthMiddleware(authService services.AuthService, userType string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := strings.Split(ctx.GetReqHeaders()["Authorization"], " ")
		if len(authHeader) == 2 && authHeader[0] == "Bearer" {
			payload, err := authService.ValidateAccessToken(authHeader[1], userType)
			if err != nil {
				ctx.Status(http.StatusUnauthorized)
				return ctx.JSON(err)
			}
			id, err := strconv.Atoi(payload.Id)
			if err != nil {
				ctx.Status(http.StatusBadRequest)
				return ctx.JSON(fiber.Map{
					"error": "wrong token payload",
				})
			}
			ctx.Locals("user-id", id)
			return ctx.Next()
		}

		ctx.Status(http.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}
}
