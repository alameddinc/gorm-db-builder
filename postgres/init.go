package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "host=127.0.0.1 user=alameddinDB password=password123 DB.name=alameddindb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection Error")
	}
	return db
}
