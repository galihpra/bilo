package repository

import (
	"bilo/features/charts"
	"context"

	"gorm.io/gorm"
)

type Chart struct {
	Id    string  `gorm:"column:id; primaryKey;"`
	Total float64 `gorm:"column:total; type:decimal(16,2);"`

	UserId uint `gorm:"column:user_id;"`
	User   User `gorm:"foreignKey:UserId"`

	Details []ChartDetail
}

type ChartDetail struct {
	Quantity int     `gorm:"column:quantity; type:integer;"`
	Subtotal float64 `gorm:"column:subtotal; type:decimal(16,2);"`

	ChartId string `gorm:"column:chart_id;"`
	Chart   Chart  `gorm:"foreignKey:ChartId;"`

	ProductId uint    `gorm:"column:product_id;"`
	Product   Product `gorm:"foreignKey:ProductId;"`
}

type User struct {
	Id uint
}

type Product struct {
	Id           uint
	Name         string
	Price        float64
	ThumbnailUrl string

	UserId uint `gorm:"column:user_id;"`
	User   User `gorm:"foreignKey:UserId"`
}

type chartRepository struct {
	db *gorm.DB
}

func NewChartRepository(db *gorm.DB) charts.Repository {
	return &chartRepository{
		db: db,
	}
}

func (repo *chartRepository) Create(ctx context.Context, data charts.Chart) error {
	var inputDB = new(Chart)

	inputDB.UserId = data.UserId

	return nil
}
