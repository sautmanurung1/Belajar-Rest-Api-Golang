package database

import (
	"REST-api/config"
	"REST-api/models"
)

func GetUser() (interface{}, error){
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil{
		return nil, e
	}
	return users, nil
}

