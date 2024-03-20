package repository

import (
	"bilo/features/carts"
	"context"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Id    string  `gorm:"column:id; primaryKey;"`
	Total float64 `gorm:"column:total; type:decimal(16,2);"`

	UserId uint `gorm:"column:user_id;"`
	User   User `gorm:"foreignKey:UserId"`

	Details []CartDetail
}

type CartDetail struct {
	Quantity int     `gorm:"column:quantity; type:integer;"`
	Subtotal float64 `gorm:"column:subtotal; type:decimal(16,2);"`

	ChartId string `gorm:"column:chart_id;"`
	Chart   Cart   `gorm:"foreignKey:ChartId;"`

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

func NewChartRepository(db *gorm.DB) carts.Repository {
	return &chartRepository{
		db: db,
	}
}

func (repo *chartRepository) Create(ctx context.Context, data carts.CartDetail) error {
	var product Product
	var inputChartDetail = new(CartDetail)

	var chartId = time.Now().String()
	var total float64

	if err := repo.db.Where("id = ?", data.Products.ID).First(&product).Error; err != nil {
		return err
	}

	var subtotal = product.Price * float64(data.Quantity)
	total += subtotal

	inputChartDetail.ChartId = chartId
	inputChartDetail.ProductId = data.Products.ID
	inputChartDetail.Quantity = data.Quantity
	inputChartDetail.Subtotal = subtotal

	if err := repo.db.Create(inputChartDetail).Error; err != nil {
		return err
	}

	var inputChart = new(Cart)

	inputChart.Id = chartId
	inputChart.Total = total

	if err := repo.db.Create(inputChart).Error; err != nil {
		return err
	}

	return nil
}
