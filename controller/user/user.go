package user

import "github.com/gofiber/fiber/v2"

func GetUser(ctx *fiber.Ctx) error {
	res := Response{
		Id:   1,
		Name: "test_user",
	}
	return ctx.JSON(res)
}
