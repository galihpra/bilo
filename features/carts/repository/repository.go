package repository

import (
	"bilo/features/carts"
	"context"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Id       string  `gorm:"column:id; primaryKey;"`
	Quantity int     `gorm:"column:quantity; type:integer;"`
	Subtotal float64 `gorm:"column:subtotal; type:decimal(16,2);"`

	UserId uint `gorm:"column:user_id;"`
	User   User `gorm:"foreignKey:UserId"`

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

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) carts.Repository {
	return &cartRepository{
		db: db,
	}
}

func (repo *cartRepository) Create(ctx context.Context, data carts.Cart) error {
	var product Product

	var cartId = time.Now().String()
	var subtotal float64

	if err := repo.db.Where("id = ?", data.ProductId).First(&product).Error; err != nil {
		return err
	}
	subtotal = product.Price * float64(data.Quantity)

	var inputChart = new(Cart)

	inputChart.Id = cartId
	inputChart.Subtotal = subtotal
	inputChart.ProductId = data.ProductId
	inputChart.UserId = data.UserId
	inputChart.Quantity = data.Quantity

	if err := repo.db.Create(inputChart).Error; err != nil {
		return err
	}

	return nil
}
