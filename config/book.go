package config

import (
	"REST-api/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbBook *gorm.DB
func InitDBBook(){
	config := map[string]string{
		"DbBook_Username" : "root",
		"DbBook_Password" : "Sautmanurung234",
		"DbBook_Port" : "3306",
		"DbBook_Host" : "127.0.0.1",
		"DbBook_Name" : "books",
	}
	connectionString := 
	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	config["DbBook_Username"],
	config["DbBook_Password"],
	config["DbBook_Host"],
	config["DbBook_Port"],
	config["DbBook_Name"])
	var e error
	DbBook, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil{
		panic(e.Error())
	}
	initMigrate()
}

func initMigrate(){
	DbBook.AutoMigrate(&models.Book{})
}