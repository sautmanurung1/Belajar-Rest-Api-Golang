package models

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	ID int `json:"id"  form:"id"`
	Name string `json:"name"  form:"name"`
	Email string `json:"email"  form:"email"`
	Password string `json:"password"  form:"password"`
}

