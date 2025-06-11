package handler

import "github.com/gofiber/fiber/v2"

func ReadUser(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"user": "user satu",
	})
}