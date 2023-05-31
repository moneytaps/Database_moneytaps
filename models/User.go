package models

import "gorm.io/gorm"

type User struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	SurName     string `json:"sur_name"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
