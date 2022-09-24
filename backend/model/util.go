package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MYSQL_USER     = "golang"
	MYSQL_PASSWORD = "golang"
	MYSQL_HOST     = "database"
	MYSQL_PORT     = "3306"
	MYSQL_DATABASE = "bookshelf"
	DSN            = MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
