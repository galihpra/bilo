package handler

import (
	"bilo/config"
	"bilo/features/carts"
	"bilo/helper/tokens"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type chartHandler struct {
	service   carts.Service
	jwtConfig config.JWT
}

func NewCartHandler(service carts.Service, jwtConfig config.JWT) carts.Handler {
	return &chartHandler{
		service:   service,
		jwtConfig: jwtConfig,
	}
}

func (hdl *chartHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request = new(CreateRequest)
		var response = make(map[string]any)

		token := c.Get("user")
		if token == nil {
			response["message"] = "unauthorized access"
			return c.JSON(http.StatusUnauthorized, response)
		}

		userId, err := tokens.ExtractToken(hdl.jwtConfig.Secret, token.(*jwt.Token))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "incorect input data"
			return c.JSON(http.StatusBadRequest, response)
		}

		var parseInput = new(carts.Cart)
		parseInput.UserId = userId
		parseInput.ProductId = request.ProductId
		parseInput.Quantity = request.Qty

		if err := hdl.service.Create(c.Request().Context(), *parseInput); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "unauthorized") {
				response["message"] = "unauthorized"
				return c.JSON(http.StatusBadRequest, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "product added to cart"
		return c.JSON(http.StatusCreated, response)
	}
}
