package models

import "gorm.io/gorm"

type Client struct {
	ID          uint   `gorm:"primary key; autoIncrement" json:"id"`
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

func Migrateclient(db *gorm.DB) error {
	err := db.AutoMigrate(&Client{})
	return err
}
