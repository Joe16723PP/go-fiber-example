package server

import (
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func InitApp() *fiber.App {
	app := fiber.New()
	return app
}
