package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection(ConnectionOptions map[string]interface{}) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		ConnectionOptions["user"],
		ConnectionOptions["password"],
		ConnectionOptions["host"],
		ConnectionOptions["port"],
		ConnectionOptions["db"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection Lost!")
	}
	return db
}
