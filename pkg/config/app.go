package config

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const dsn = "Bruce:palebluedot4@tcp(host.docker.internal:3306)/bookstoreAPI?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() {
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db = mysqlDB
}

func GetDB() *gorm.DB {
	return db
}
