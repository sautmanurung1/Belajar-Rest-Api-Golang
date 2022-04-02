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
	e.POST("/users", c.CreateUserController)

	// Route Books
	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)

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