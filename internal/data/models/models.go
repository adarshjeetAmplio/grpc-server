package models

import "gorm.io/gorm"


type User struct{
	gorm.Model
	Name string `json:"name"`
	Email string `gorm:"unique; not null" json:"email"`
	Password string `json:"-"`
}

func (u User) TableName() string{
	return "user";
}