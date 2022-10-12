package model_test

import (
	"testing"

	"github.com/SuperTikuwa/book_app/model"
	"github.com/brianvoe/gofakeit/v6"
)

var faker *gofakeit.Faker

func init() {
	faker = gofakeit.New(0)
}

func TestUser_Create(t *testing.T) {
	u := model.User{
		Name:        faker.Name(),
		Email:       faker.Email(),
		CognitoUUID: faker.UUID(),
	}

	if err := u.Create(); err != nil {
		t.Error(err)
	}
}
