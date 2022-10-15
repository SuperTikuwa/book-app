package model_test

import (
	"testing"
	"time"

	"github.com/SuperTikuwa/book_app/model"
	"github.com/brianvoe/gofakeit/v6"
)

var faker *gofakeit.Faker

func init() {
	faker = gofakeit.New(0)
}

func TestUser_Create(t *testing.T) {
	nowString := time.Now().String()
	type fields struct {
		Name        string
		Email       string
		CognitoUUID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "正常系データ",
			fields: fields{
				Name:        "test" + nowString,
				Email:       faker.Email(),
				CognitoUUID: faker.UUID(),
			},
			wantErr: false,
		},
		{
			name: "異常系-Nameが空",
			fields: fields{
				Email:       faker.Email(),
				CognitoUUID: faker.UUID(),
			},
			wantErr: true,
		},
		{
			name: "異常系-Emailが空",
			fields: fields{
				Name:        faker.Name(),
				CognitoUUID: faker.UUID(),
			},
			wantErr: true,
		},
		{
			name: "異常系-Nameが重複",
			fields: fields{
				Name:        "test" + nowString,
				Email:       faker.Email(),
				CognitoUUID: faker.UUID(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := model.User{
				Name:        tt.fields.Name,
				Email:       tt.fields.Email,
				CognitoUUID: tt.fields.CognitoUUID,
			}
			if err := u.Create(); (err != nil) != tt.wantErr {
				t.Errorf("User.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
