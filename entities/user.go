package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"unique" json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	ReferalCode string `json:"referal_code" form:"referal_code"`
	Reference   string `json:"reference" form:"reference"`
}
