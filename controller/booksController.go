package controller

import (
	"REST-api/config"
	"REST-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var(
	books []models.Book
)

// get all users
func GetBooksController(c echo.Context) error {
	if err := config.DbBook.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages" : "Success Get All Books",
		"users" : books,
	})
}

// get user by id
func GetBookController(c echo.Context) error{
	// get user by id
	ID , _ := strconv.Atoi(c.Param("id"))
	for _, book := range books{
		if book.ID == ID{
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message" : "Success Get Books",
				"User" : book,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message" : "Books Not Found In Databases",
	})
}

func CreateBooksController(c echo.Context) error{
	book := models.Book{}
	c.Bind(&book)
	if err := config.DbBook.Save(&book).Error; err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Success Create New Books",
		"User" : book,
	})
}

// delete user by id 
func DeleteBooksController(c echo.Context) error{
	// delete user by id
	ID , _ := strconv.Atoi(c.Param("id"))
	for index, book := range books{
		if book.ID == ID{
			users = append(users[:index], users[index+1:]...)
			if err := config.DbBook.Delete(&book).Error; err != nil{
				return echo.NewHTTPError(http.StatusBadRequest,err.Error())
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message" : "Success Delete Books",
				"User" : book,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message" : "Books Not Found In Databases",
	})
}

// Update user by id
func UpdateBooksController(c echo.Context) error{
	// update user by id
	ID , _ := strconv.Atoi(c.Param("id"))
	for index, book := range books{
		if book.ID == ID{
			c.Bind(&book)
			books[index] = book
			if err := config.DbBook.Save(&book).Error; err != nil{
				return echo.NewHTTPError(http.StatusBadRequest,err.Error())
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message" : "Success Update Books",
				"User" : book,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message" : "Books Not Found In Databases",
	})
}