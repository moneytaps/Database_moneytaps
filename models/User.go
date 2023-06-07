package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint   `gorm:"primary key;autoIncrement" json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Contact    string `json:"contact"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	SurName    string `json:"surname"`
}
type LoanStatus struct {
	ID          uint    `gorm:"primary key;autoIncrement" json:"id"`
	Status      string  `json:"status"`
	Contact     string  `json:"contact"`
	FirstName   string  `json:"first_name"`
	MiddleName  string  `json:"middle_name"`
	SurName     string  `json:"sur_name"`
	Birth       string  `json:"birth"`
	Gender      string  `json:"gender"`
	Address     string  `json:"address"`
	Primary     string  `json:"primary"`
	LoanAmount  float64 `json:"loan_amount"`
	Days        string  `json:"days"`
	Interest    float64 `json:"interest"`
	TotalAmount float64 `json:"total_amount"`
}
type History struct {
	ID          uint    `gorm:"primary key;autoIncrement" json:"id"`
	Status      string  `json:"status"`
	Contact     string  `json:"contact"`
	FirstName   string  `json:"first_name"`
	MiddleName  string  `json:"middle_name"`
	SurName     string  `json:"sur_name"`
	Birth       string  `json:"birth"`
	Gender      string  `json:"gender"`
	Address     string  `json:"address"`
	Primary     string  `json:"primary"`
	LoanAmount  float64 `json:"loan_amount"`
	Days        int     `json:"days"`
	Interest    float64 `json:"interest"`
	TotalAmount float64 `json:"total_amount"`
	CreatedAt   time.Time
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &LoanStatus{}, &History{})
	return err
}
