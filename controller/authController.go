package controller

import (
	"REST-api/config"
	"REST-api/middleware"
	"REST-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error{
	user := models.User{}

	c.Bind(&user)

	err := config.DbUser.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "Email or Password is wrong",
		})
	}
	token, err := middleware.CreateToken(strconv.Itoa(int(user.ID)), user.Name)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "Error Create Token",
		})
	}
	
	responseData := models.Auth{}
	responseData.ID = user.Id
	responseData.Name = user.Name
	responseData.Email = user.Email
	responseData.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Success Login",
		"Respone Data" : responseData,
		"Token" : token,
	})
}