package repository

import (
	"bilo/features/charts"
	"context"
	"time"

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

func (repo *chartRepository) Create(ctx context.Context, data []charts.ChartDetail) error {
	var product Product
	var inputChartDetail = new(ChartDetail)

	var chartId = time.Now().String()
	var total float64

	for i := 0; i < len(data); i++ {
		if err := repo.db.Where("id = ?", data[i].Products.ID).First(&product).Error; err != nil {
			return err
		}

		var subtotal = product.Price * float64(data[i].Quantity)
		total += subtotal

		inputChartDetail.ChartId = chartId
		inputChartDetail.ProductId = data[i].Products.ID
		inputChartDetail.Quantity = data[i].Quantity
		inputChartDetail.Subtotal = subtotal

		if err := repo.db.Create(inputChartDetail).Error; err != nil {
			return err
		}
	}

	var inputChart = new(Chart)

	inputChart.Id = chartId
	inputChart.Total = total

	if err := repo.db.Create(inputChart).Error; err != nil {
		return err
	}

	return nil
}
