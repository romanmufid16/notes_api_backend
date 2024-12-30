package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/notes_api_backend/pkg/utils"
)

func ErrorMiddleware(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	errorResponse := utils.BuildErrorResponse(err.Error())
	return ctx.Status(code).JSON(errorResponse)
}
