package routes

import (
	"app/constants"
	"app/controllers"
	"app/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddleware(e)
	e.GET("/users", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)

	eBasic := e.Group("/product")
	eBasic.Use(middleware.BasicAuth(middlewares.BasicAuthDB))
	eBasic.GET("/", controllers.GetUserControllers)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.KEY_JWT)))
	eJwt.GET("/users", controllers.GetUserControllers)
	return e
}
