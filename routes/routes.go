package routes

import (
	"REST-api/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo{
	e := echo.New()
	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	// Route Books
	e.GET("/books", controller.GetBooksController)
	e.POST("/books", controller.CreateBooksController)
	e.DELETE("/books/:id", controller.DeleteBooksController)
	e.PUT("/books/:id", controller.UpdateBooksController)
	return e
}