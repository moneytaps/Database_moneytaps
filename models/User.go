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
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	SurName    string `json:"sur_name"`
}
type LoanStatus struct {
	ID          uint   `gorm:"primary key; autoIncrement" json:"id"`
	Status      int    `json:"status"`
	Contact     string `json:"contact"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	SurName     string `json:"sur_name"`
	Birth       string `json:"birth"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Primary     string `json:"primary"`
	LoanAmount  int    `json:"loan_amount"`
	Days        int    `json:"days"`
	Interest    int    `json:"interest"`
	TotalAmount int    `json:"total_amount"`
	Purpose     string `json:"purpose"`
}
type History struct {
	ID          uint   `gorm:"primary key; autoIncrement" json:"id" `
	Contact     string `json:"contact"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	SurName     string `json:"sur_name"`
	Birth       string `json:"birth"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Primary     string `json:"primary"`
	LoanAmount  int    `json:"loan_amount"`
	Days        int    `json:"days"`
	Interest    int    `json:"interest"`
	TotalAmount int    `json:"total_amount"`
	Purpose     string `json:"purpose"`
	CreatedAt   time.Time
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &LoanStatus{}, &History{})
	return err
}
