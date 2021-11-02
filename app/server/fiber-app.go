package server

import (
	"github.com/gofiber/fiber/v2"
)

func InitApp() *fiber.App {
	app := fiber.New()
	return app
}
