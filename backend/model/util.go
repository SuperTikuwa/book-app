package model

import (
	"database/sql"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

const (
	MYSQL_USER     = "golang"
	MYSQL_PASSWORD = "golang"
	MYSQL_DATABASE = "bookshelf"
)

var (
	MYSQL_HOST = "database"
	MYSQL_PORT = "3306"
)

func dsn(isTesting bool) string {
	if isTesting {
		MYSQL_HOST = "localhost"
		MYSQL_PORT = "33062"
	}

	return MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func Connect() *bun.DB {

	engine, err := sql.Open("mysql", dsn(flag.Lookup("test.v") != nil))
	if err != nil {
		panic(err)
	}
	return bun.NewDB(engine, mysqldialect.New())
}
