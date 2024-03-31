package service

import (
	"bilo/features/carts"
	"context"
	"errors"
)

type cartService struct {
	repo carts.Repository
}

func NewCartService(repo carts.Repository) carts.Service {
	return &cartService{
		repo: repo,
	}
}

func (srv *cartService) Create(ctx context.Context, data carts.Cart) error {
	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv *cartService) GetByUserId(ctx context.Context, UserId uint) ([]carts.Cart, error) {
	if UserId == 0 {
		return nil, errors.New("validate: invalid user id")
	}

	result, err := srv.repo.GetByUserId(ctx, UserId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
