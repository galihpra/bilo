package handler

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Hp       string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
	Image    string `json:"user_image" form:"user_image"`
}
