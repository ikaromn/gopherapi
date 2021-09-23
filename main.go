package main

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/ikaromn/go-api/gopherapi"
)

func main() {
	app := fiber.New()
	routes.Route(app)
	app.Listen(":3001")
}
