package service

import (
	"bilo/features/products"
	"context"
	"errors"
)

type productService struct {
	repo products.Repository
}

func NewProductService(repo products.Repository) products.Service {
	return &productService{
		repo: repo,
	}
}

func (srv *productService) Create(ctx context.Context, data products.Product) error {
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

	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
