package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/controller/healthCheck"
)

func RegisterHealthCheckRouter(router *fiber.Router) {
	r := *router // router pointer
	r.Get("/", healthCheck.GetHealthCheck)
}
