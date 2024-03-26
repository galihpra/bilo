package handler

type CartResponse struct {
	Id       string  `json:"cart_id"`
	Qty      int     `json:"qty"`
	Subtotal float64 `json:"subtotal"`

	Product ProductResponse `json:"products"`
}

type ProductResponse struct {
	Id        uint    `json:"product_id"`
	Name      string  `json:"product_name"`
	Thumbnail string  `json:"thumbnail"`
	Price     float64 `json:"price"`

	Users UserResponse `json:"seller"`
}

type UserResponse struct {
	Id       uint   `json:"user_id"`
	Username string `json:"username"`
}
