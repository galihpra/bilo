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
}

type Handler interface {
	Create() echo.HandlerFunc
}

type Service interface {
	Create(ctx context.Context, data Cart) error
}

type Repository interface {
	Create(ctx context.Context, data Cart) error
}
