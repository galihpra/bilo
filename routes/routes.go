package routes

import (
	"bilo/features/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uh users.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uh)
}

func routeUser(e *echo.Echo, uh users.Handler) {
	e.POST("/register", uh.Register())
	e.POST("/login", uh.Login())
}
