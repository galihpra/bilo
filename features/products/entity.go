package products

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Product struct {
	ID          uint
	Name        string
	Rating      float32
	Price       float64
	Stock       int64
	Condition   string
	Description string
}

type Handler interface {
	Create() echo.HandlerFunc
}

type Service interface {
	Create(token *jwt.Token, data Product) error
}

type Repository interface {
	Create(userID uint, data Product) error
}
