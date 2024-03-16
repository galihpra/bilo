package handler

type CreateRequest struct {
	UserId    uint
	ProductId []Product
}

type Product struct {
	ProductId uint `json:"product_id" form:"product_id"`
	Qty       uint `json:"qty" form:"qty"`
}
