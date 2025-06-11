package route

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App)  {
	route.Get("/user", handler.UserFindAll)
	route.Get("/user/:id", handler.UserFindById)
	route.Post("/user", handler.UserCreate)
}