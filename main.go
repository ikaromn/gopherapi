package main

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/ikaromn/go-api/costumers"
)

func main() {
	app := fiber.New()
	routes.Route(app)
	app.Listen(":3001")
}
