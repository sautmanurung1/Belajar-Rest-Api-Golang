package config

import (
	"REST-api/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbUser *gorm.DB

func InitDBUser(){
	config := map[string]string{
		"DbUser_Username" : "root",
		"DbUser_Password" : "Sautmanurung234",
		"DbUser_Port" : "3306",
		"DbUser_Host" : "127.0.0.1",
		"DbUser_Name" : "user",
	}
	connectionString := 
	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	config["DbUser_Username"],
	config["DbUser_Password"],
	config["DbUser_Host"],
	config["DbUser_Port"],
	config["DbUser_Name"])
	var e error
	DbUser, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil{
		panic(e.Error())
	}
	initmigrate()
}

func initmigrate(){
	DbUser.AutoMigrate(&models.User{})
}