package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Name string `json:"name", form:"name"`
	Email string `json:"email", form:"email"`
	Token string `json:"token", form:"token"`
}