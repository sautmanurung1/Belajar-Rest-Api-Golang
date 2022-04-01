package models

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Id uint `json:"uid", form:"uid"`
	ID int `json:"id"  form:"id"`
	Name string `json:"name"  form:"name"`
	Email string `json:"email"  form:"email"`
	Password string `json:"password"  form:"password"`
}

