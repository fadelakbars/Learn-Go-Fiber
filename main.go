package main

import (
	"go-fiber/database"
	"go-fiber/database/migration"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INISIALISASI DATABASE
	database.InitDatabase()
	// MIGRASI DATABASE
	migration.RunMigration()

	app := fiber.New()

	// INISIALISASI ROUTING
	route.RouteInit(app)

	app.Listen(":3000")
}