package gopherapi

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/ikaromn/go-api/settings"
)

type View struct {
	model       interface{}
	models      interface{}
	lookupField string
}

func (v View) GetAndListView(ctx *fiber.Ctx) error {
	db := settings.DbOpenConnection()
	defer settings.DbCloseConnection(db)

	if ctx.Params(v.lookupField) != "" {
		lookupField := ctx.Params(v.lookupField)
		statemant := v.lookupField + " = ?"

		model := reflect.TypeOf(v.model)
		m := reflect.New(model).Interface()

		db.Where(statemant, lookupField).Find(m)
		return ctx.JSON(m)
	}

	models := reflect.TypeOf(v.models)
	m := reflect.New(models).Interface()

	result := db.Find(m)
	ApiListResult(ctx, m, result.RowsAffected)

	return nil
}

func (v View) CreateView(ctx *fiber.Ctx) error {
	db := settings.DbOpenConnection()
	defer settings.DbCloseConnection(db)

	model := reflect.TypeOf(v.model)
	m := reflect.New(model).Interface()
	if err := ctx.BodyParser(m); err != nil {
		return err
	}

	if err := db.Create(m).Error; err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"error": err,
		})
	}
	ctx.JSON(m)

	return nil
}

func (v View) PartialUpdateView(ctx *fiber.Ctx) error {
	db := settings.DbOpenConnection()
	defer settings.DbCloseConnection(db)

	lookupField := ctx.Params(v.lookupField)
	statemant := v.lookupField + " = ?"

	model := reflect.TypeOf(v.model)
	m := reflect.New(model).Interface()

	if err := ctx.BodyParser(m); err != nil {
		return err
	}

	if err := db.Where(statemant, lookupField).Updates(m).Error; err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"error": err,
		})
	}

	db.Where(statemant, lookupField).Find(m)
	return ctx.JSON(m)
}

var costumer Costumer
var costumers []Costumer

var CostumersView = View{
	model:       costumer,
	models:      costumers,
	lookupField: "document_id",
}
