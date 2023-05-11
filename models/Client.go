package models

import "gorm.io/gorm"

type Client struct {
	First_Name  string `json:"firstname"`
	Middle_Name string `json:"middlename"`
	Sur_Name    string `json:"surname"`
	Birthdate   string `json:"birthdate"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Primary     string `json:"primary"`
	LoanAmount  int    `json:"loanAmount"`
	LoanDays    int    `json:"loanDays"`
	Interest    int    `json:"interest"`
	TotalAmount int    `json:"totalAmount"`
	Purpose     string `json:"porpose"`
}

func Migrateclient(db *gorm.DB) error {
	err := db.AutoMigrate(&Client{})
	return err
}
