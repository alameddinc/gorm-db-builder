package postgres

import (
	"fmt"
	"github.com/alameddinc/gorm-db-builder/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	ConnectionOptions := configs.GetConfigs("postgres")
	dsn := fmt.Sprintf("host=%v user=%v password=%v DB.name=%v port=%d sslmode=%v",
		ConnectionOptions["host"],
		ConnectionOptions["username"],
		ConnectionOptions["password"],
		ConnectionOptions["db"],
		ConnectionOptions["port"],
		ConnectionOptions["ssl"])
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection Error")
	}
	return db
}
