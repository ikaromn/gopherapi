package gopherapi

import "github.com/gofiber/fiber/v2"

func ApiListResult(ctx *fiber.Ctx, result interface{}, total int64) {
	ctx.JSON(fiber.Map{
		"total": total,
		"data":  result,
	})
}
