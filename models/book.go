package models

import (
	"gorm.io/gorm"
)

type Book struct{
	gorm.Model
	ID int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Author string `json:"author" form:"author"`
	Year string `json:"year" form:"year"`
}