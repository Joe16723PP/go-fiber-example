package healthCheck

import "github.com/gofiber/fiber/v2"

func GetHealthCheck(ctx *fiber.Ctx) error {
	res := Response{
		Version: "v1.0",
		Status:  "still alive",
	}
	return ctx.JSON(res)
}
