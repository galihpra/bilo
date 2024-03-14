package charts

import (
	"context"
	"io"

	"github.com/labstack/echo/v4"
)

type Chart struct {
	ID    string
	Total float64

	UserId  uint
	Details []ChartDetail
}

type ChartDetail struct {
	Quantity int
	Subtotal float64

	ChartId  string
	Products Product
}

type Product struct {
	ID        uint
	Name      string
	Price     float64
	Thumbnail Image

	UserId uint
}

type Image struct {
	ID       uint
	ImageURL string
	ImageRaw io.Reader
}

type Handler interface {
	Create() echo.HandlerFunc
}

type Service interface {
	Create(ctx context.Context, data Chart) error
}

type Repository interface {
	Create(ctx context.Context, data Chart) error
}
