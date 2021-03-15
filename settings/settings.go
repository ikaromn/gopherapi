package settings

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbOpenConnection() *gorm.DB {
	dns_str := "postgresql://go_api:go_api@localhost:5432/go_api"
	db, _ := gorm.Open(postgres.Open(dns_str), &gorm.Config{})
	return db
}

func DbCloseConnection(db *gorm.DB) {
	dbconn, _ := db.DB()
	dbconn.Close()
}
