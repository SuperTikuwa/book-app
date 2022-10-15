package handler_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
)

var faker *gofakeit.Faker

func init() {
	faker = gofakeit.New(0)
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}
}
