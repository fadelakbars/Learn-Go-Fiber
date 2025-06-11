package main

import (

	"github.com/gofiber/fiber/v2"
	"go-fiber/route"
)

func main() {
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}