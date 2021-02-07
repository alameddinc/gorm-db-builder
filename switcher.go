package gormDBPlus

import (
	"github.com/alameddinc/gorm-db-builder/mysql"
	"github.com/alameddinc/gorm-db-builder/postgres"
	gorm "gorm.io/gorm"
)

const (
	POSTGRES_DRIVE = "postgre"
	MYSQL_DRIVE    = "mysql"
)

var connectionList = map[string]func() *gorm.DB{
	POSTGRES_DRIVE: postgres.Connection,
	MYSQL_DRIVE:    mysql.Connection,
}
