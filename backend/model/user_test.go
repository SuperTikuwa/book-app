package model_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SuperTikuwa/book_app/model"
	"github.com/brianvoe/gofakeit"
)

func TestUser_Create(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err, "mock failed")
	}
	d, _ := db.DB()
	defer d.Close()

	u := model.User{
		Name:        gofakeit.Name(),
		Email:       gofakeit.Email(),
		CognitoUUID: gofakeit.UUID(),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`cognito_uuid`) VALUES (?,?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	if err := u.Create(db); err != nil {
		t.Fatal(err)
	}
}
