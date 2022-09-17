package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DSN = "golang:golang@tcp(database:3306)/bookshelf?charset=utf8mb4&parseTime=True&loc=Local"

func connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
