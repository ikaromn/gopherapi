package costumers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikaromn/go-api/settings"
)

func GetAndListCostumer(ctx *fiber.Ctx) error {
	db := settings.DbOpenConnection()
	defer settings.DbCloseConnection(db)

	if ctx.Params("document_id") != "" {
		var costumer Costumer
		documentId := ctx.Params("document_id")
		db.Where("document_id = ?", documentId).Find(&costumer)
		return ctx.JSON(costumer)
	}
	var costumers []Costumer
	result := db.Find(&costumers)
	ApiListResult(ctx, costumers, result.RowsAffected)
	return nil
}

func CreateAndUpdateCostumer(ctx *fiber.Ctx) error {
	db := settings.DbOpenConnection()
	defer settings.DbCloseConnection(db)

	var costumer Costumer

	if err := ctx.BodyParser(&costumer); err != nil {
		return err
	}

	db.Create(&costumer)

	ctx.JSON(costumer)

	return nil
}
