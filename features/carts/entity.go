package carts

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Cart struct {
	ID       string
	Quantity int
	Subtotal float64

	UserId    uint
	ProductId uint

	Product Product
}

type Product struct {
	ID        uint
	Name      string
	Price     float64
	Thumbnail string

	User User
}

type User struct {
	ID       uint
	Username string
}

type Handler interface {
	Create() echo.HandlerFunc
	GetByUserId() echo.HandlerFunc
}

type Service interface {
	Create(ctx context.Context, data Cart) error
	GetByUserId(ctx context.Context, UserId uint) ([]Cart, error)
}

type Repository interface {
	Create(ctx context.Context, data Cart) error
	GetByUserId(ctx context.Context, UserId uint) ([]Cart, error)
}
