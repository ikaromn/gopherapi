package main

import (
	models "github.com/ikaromn/go-api/costumers"
	"github.com/ikaromn/go-api/settings"
)

func main() {
	db := settings.DbOpenConnection()
	defer settings.DbCloseConnection(db)

	db.Migrator().CreateTable(&models.Costumer{})
}
