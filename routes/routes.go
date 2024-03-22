package routes

import (
	"bilo/features/carts"
	"bilo/features/products"
	"bilo/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	JWTKey         string
	Server         *echo.Echo
	UserHandler    users.Handler
	ProductHandler products.Handler
	CartHandler    carts.Handler
}

func (router Routes) InitRouter() {
	router.UserRouter()
	router.ProductRouter()
	router.CartRouter()
}

func (router *Routes) UserRouter() {
	router.Server.POST("/register", router.UserHandler.Register())
	router.Server.POST("/login", router.UserHandler.Login())
}

func (router *Routes) ProductRouter() {
	router.Server.POST("/products", router.ProductHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))
}

func (router *Routes) CartRouter() {
	router.Server.POST("/carts", router.CartHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))
}
