package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Email       string
	CognitoUUID string
}

func (u User) Create() error {
	db := connect()
	d, _ := db.DB()
	defer d.Close()

	if db.Create(&u).Error != nil {
		return db.Create(&u).Error
	}

	return nil
}
