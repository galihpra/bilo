package products

import (
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

	UserId uint
}

type Handler interface {
	Create() echo.HandlerFunc
}

type Service interface {
	Create(data Product) error
}

type Repository interface {
	Create(data Product) error
}
