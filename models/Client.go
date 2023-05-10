package models

import "gorm.io/gorm"

type Client struct {
	First_Name  string `json:"firstname"`
	Middle_Name string `json:"middlename"`
	Sur_Name    string `json:"surname"`
	Birthdate   int16  `json:"birthdate"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Primary     string `json:"primary "`
	LoanAmount  int16  `json:"loanAmount"`
	LoanDays    int16  `json:"loanDays"`
	Interest    int16  `json:"interest"`
	TotalAmount int16  `json:"totalAmount"`
	Purpose     string `json:"porpose"`
}

func Migrateclient(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
