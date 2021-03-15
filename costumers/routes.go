package costumers

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/costumers/:document_id?", GetAndListCostumer)
	app.Post("/costumers/", CreateAndUpdateCostumer)
}
