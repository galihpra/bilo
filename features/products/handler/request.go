package handler

import "io"

type CreateRequest struct {
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Stock       int64   `json:"stock" form:"stock"`
	Condition   string  `json:"condition" form:"condition"`
	Price       float64 `json:"price" form:"price"`
	UserId      uint
	Images      []io.Reader
}
