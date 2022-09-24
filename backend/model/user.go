package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Email       string
	CognitoUUID string
}

func (u User) Create(db *gorm.DB) error {
	d, _ := db.DB()
	defer d.Close()

	// db.Begin()
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	// db.Commit()

	return nil
}
