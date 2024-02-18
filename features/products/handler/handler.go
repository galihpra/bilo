package handler

import (
	"bilo/config"
	"bilo/features/products"
	"bilo/helper/tokens"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	service   products.Service
	jwtConfig config.JWT
}

func NewProductHandler(service products.Service, jwtConfig config.JWT) products.Handler {
	return &productHandler{
		service:   service,
		jwtConfig: jwtConfig,
	}
}

func (hdl *productHandler) Create() echo.HandlerFunc {
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

		if err := c.Request().ParseMultipartForm(10 << 20); err != nil {
			c.Logger().Error(err)
			response["message"] = "failed to parse form data"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "incorect input data"
			return c.JSON(http.StatusBadRequest, response)
		}

		var parseInput = new(products.Product)
		parseInput.Name = request.Name
		parseInput.Price = request.Price
		parseInput.Stock = request.Stock
		parseInput.Condition = request.Condition
		parseInput.Description = request.Description
		parseInput.UserId = userId

		// Handle file uploads
		images := c.Request().MultipartForm.File["images"]
		var imageReaders []io.Reader
		for _, img := range images {
			file, err := img.Open()
			if err != nil {
				c.Logger().Error(err)
				response["message"] = "failed to open uploaded file"
				return c.JSON(http.StatusInternalServerError, response)
			}
			imageReaders = append(imageReaders, file) // Tambahkan file ke dalam slice sebelum menutupnya
			defer file.Close()
		}
		// Set image readers to request
		request.Images = imageReaders
		fmt.Println(request.Images)

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

		response["message"] = "create product success"
		return c.JSON(http.StatusCreated, response)
	}
}
