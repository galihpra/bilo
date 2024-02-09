package repository

import (
	"bilo/features/products"

	"gorm.io/gorm"
)

type Product struct {
	Id          uint    `gorm:"column:id; primaryKey;"`
	Name        string  `gorm:"column:name; type:varchar(200);"`
	Rating      float32 `gorm:"column:rating; type:float"`
	Price       float64 `gorm:"column:price; type:decimal(16,2);"`
	Stock       int64   `gorm:"column:stock; type:integer;"`
	Condition   string  `gorm:"column:condition; type:varchar(8);"`
	Description string  `gorm:"column:description; type:text;"`

	UserId uint `gorm:"column:user_id;"`
	User   User `gorm:"foreignKey:UserId"`
}

type User struct {
	Id uint
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) products.Repository {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) Create(data products.Product) error {
	var inputDB = new(Product)
	inputDB.Name = data.Name
	inputDB.Price = data.Price
	inputDB.Stock = data.Stock
	inputDB.Condition = data.Condition
	inputDB.Description = data.Description
	inputDB.UserId = data.UserId

	if err := repo.db.Create(inputDB).Error; err != nil {
		return err
	}

	return nil
}
