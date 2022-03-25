package models

import (
	"gorm.io/gorm"
)

type Book struct{
	gorm.Model
	ID int `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Year string `json:"year" form:"year"`
}