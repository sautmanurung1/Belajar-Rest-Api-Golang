package controller

import (
	"REST-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var(
	db *gorm.DB
	users []models.User
)

// get all users
func GetUsersController(c echo.Context) error {
	if err := db.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages" : "success get all users",
		"users" : users,
	})
}

// get user by id
func GetUserController(c echo.Context) error{
	// get user by id
	ID , _ := strconv.Atoi(c.Param("id"))
	for _, user := range users{
		if user.ID == ID{
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message" : "Success Get User",
				"User" : user,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message" : "User Not Found In Databases",
	})
}

func CreateUserController(c echo.Context) error{
	user := models.User{}
	c.Bind(&user)
	if err := db.Save(&user).Error; err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Success Create New User",
		"User" : user,
	})
}

// delete user by id 
func DeleteUserController(c echo.Context) error{
	// delete user by id
	ID , _ := strconv.Atoi(c.Param("id"))
	for index, user := range users{
		if user.ID == ID{
			users = append(users[:index], users[index+1:]...)
			if err := db.Delete(&user).Error; err != nil{
				return echo.NewHTTPError(http.StatusBadRequest,err.Error())
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message" : "Success Delete User",
				"User" : user,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message" : "User Not Found In Databases",
	})
}

// Update user by id
func UpdateUserController(c echo.Context) error{
	// update user by id
	ID , _ := strconv.Atoi(c.Param("id"))
	for index, user := range users{
		if user.ID == ID{
			c.Bind(&user)
			users[index] = user
			if err := db.Save(&user).Error; err != nil{
				return echo.NewHTTPError(http.StatusBadRequest,err.Error())
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message" : "Success Update User",
				"User" : user,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message" : "User Not Found In Databases",
	})
}