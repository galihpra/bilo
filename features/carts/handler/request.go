package handler

type CreateRequest struct {
	UserId    uint
	ProductId uint `json:"product_id" form:"product_id"`
	Qty       int  `json:"qty" form:"qty"`
}