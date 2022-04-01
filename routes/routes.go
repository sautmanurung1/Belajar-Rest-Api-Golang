package routes

import (
	"REST-api/constant"
	c "REST-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo{
	e := echo.New()
	e.POST("/auth/login", c.LoginController)
	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)

	// Route Books
	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	e.POST("/books", c.CreateBooksController)
	e.DELETE("/books/:id", c.DeleteBooksController)
	e.PUT("/books/:id", c.UpdateBooksController)

	// JWT
	jwtAuth := e.Group("/jwt/redirected")
	jwtAuth.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	// Route with JWT Auth
	jwtAuth.GET("/users", c.GetUsersController)
	jwtAuth.GET("/users/:id", c.GetUserController)
	jwtAuth.DELETE("/users/:id", c.DeleteUserController)
	jwtAuth.PUT("/users/:id", c.UpdateUserController)

	jwtAuth.POST("/books", c.CreateBooksController)
	jwtAuth.DELETE("/books/:id", c.DeleteBooksController)
	jwtAuth.PUT("/books/:id", c.UpdateBooksController)
	return e
}