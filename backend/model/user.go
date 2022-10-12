package model

import (
	"context"
)

type User struct {
	Name        string
	Email       string
	CognitoUUID string
}

func (u *User) Create() error {
	db := Connect()
	defer db.Close()

	if _, err := db.NewInsert().Model(u).Exec(context.Background()); err != nil {
		return err
	}

	return nil
}
