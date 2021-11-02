package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/controller/user"
)

func RegisterUserRouter(router *fiber.Router) {
	r := *router // pointer
	r.Get("/", user.GetUser)
}
