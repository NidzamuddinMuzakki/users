package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nidzamTest.com/controller"
	"nidzamTest.com/middlewares"
)

func InitRouter(UserController controller.UserController) *echo.Echo {

	r := echo.New()
	r.Use(middleware.CORS())
	r.Use(middlewares.Auth)
	r.Use(middlewares.Recover)
	USER := r.Group("user")
	{
		USER.GET("", func(c echo.Context) error {
			UserController.FindAll(c)
			return nil
		})
		USER.GET("/byUsername", func(c echo.Context) error {
			UserController.FindByUsername(c)
			return nil
		})
		USER.PUT("/detail", func(c echo.Context) error {
			UserController.Insert(c)
			return nil
		})
		USER.POST("/detail", func(c echo.Context) error {
			UserController.Update(c)
			return nil
		})
		USER.DELETE("/detail", func(c echo.Context) error {
			UserController.Delete(c)
			return nil
		})
		USER.POST("/login", func(c echo.Context) error {
			UserController.Login(c)
			return nil
		})
	}
	return r
}
