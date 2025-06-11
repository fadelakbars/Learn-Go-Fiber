package main

import (
	"go-fiber/database"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INISIALISASI DATABASE

	database.InitDatabase()
	app := fiber.New()

	// INISIALISASI ROUTING
	route.RouteInit(app)

	app.Listen(":3000")
}