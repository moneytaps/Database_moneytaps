package models

import "gorm.io/gorm"

type User struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
