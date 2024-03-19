package service

import (
	"bilo/features/charts"
	"context"
)

type chartService struct {
	repo charts.Repository
}

func NewChartService(repo charts.Repository) charts.Service {
	return &chartService{
		repo: repo,
	}
}

func (srv *chartService) Create(ctx context.Context, data charts.CartDetail) error {
	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
