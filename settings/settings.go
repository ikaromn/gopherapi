package settings

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbOpenConnection() *gorm.DB {
	dns_str := os.Getenv("DB_DNS")
	db, _ := gorm.Open(postgres.Open(dns_str), &gorm.Config{})
	return db
}

func DbCloseConnection(db *gorm.DB) {
	dbconn, _ := db.DB()
	dbconn.Close()
}
