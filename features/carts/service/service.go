package service

import (
	"bilo/features/carts"
	"context"
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
	panic("unimplemented")
}
