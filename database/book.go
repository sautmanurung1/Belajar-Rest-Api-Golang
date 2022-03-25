package database

import (
	"REST-api/config"
	"REST-api/models"
)

func GetBook() (interface{}, error){
	var books []models.Book

	if e := config.DB.Find(&books).Error; e != nil{
		return nil, e
	}
	return books, nil
}

