package service

import (
	"bilo/config"
	"bilo/features/products"
	"bilo/helper/tokens"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type productService struct {
	repo      products.Repository
	jwtConfig config.JWT
}

func NewProductService(repo products.Repository, jwtConfig config.JWT) products.Service {
	return &productService{
		repo:      repo,
		jwtConfig: jwtConfig,
	}
}

func (srv *productService) Create(token *jwt.Token, data products.Product) error {
	if data.Name == "" {
		return errors.New("validate: name can't be empty")
	}
	if data.Price == 0 {
		return errors.New("validate: price can't be empty")
	}
	if data.Condition == "" {
		return errors.New("validate: condition can't be empty")
	}
	if data.Description == "" {
		return errors.New("validate: description can't be empty")
	}
	if data.Stock == 0 {
		return errors.New("validate: stock can't be empty")
	}

	userId, err := tokens.ExtractToken(srv.jwtConfig.Secret, token)
	if err != nil {
		return errors.New("unauthorized")
	}

	if err := srv.repo.Create(userId, data); err != nil {
		return err
	}

	return nil
}
