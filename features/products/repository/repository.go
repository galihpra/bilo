package repository

import (
	"bilo/features/products"
	"bilo/utils/cloudinary"
	"context"
	"io"

	"gorm.io/gorm"
)

type Product struct {
	Id           uint    `gorm:"column:id; primaryKey;"`
	Name         string  `gorm:"column:name; type:varchar(200);"`
	Rating       float32 `gorm:"column:rating; type:float"`
	Price        float64 `gorm:"column:price; type:decimal(16,2);"`
	Stock        int64   `gorm:"column:stock; type:integer;"`
	Condition    string  `gorm:"column:condition; type:varchar(8);"`
	Description  string  `gorm:"column:description; type:text;"`
	ThumbnailUrl string  `gorm:"column:thumbnail; type:text;"`

	UserId uint `gorm:"column:user_id;"`
	User   User `gorm:"foreignKey:UserId"`

	Images []Image `gorm:"many2many:product_images;"`
}

type User struct {
	Id uint
}

type Image struct {
	Id       uint      `gorm:"column:id; primaryKey;"`
	ImageURL string    `gorm:"column:image_url; type:text"`
	ImageRaw io.Reader `gorm:"-"`
}

type productRepository struct {
	db    *gorm.DB
	cloud cloudinary.Cloud
}

func NewProductRepository(db *gorm.DB, cloud cloudinary.Cloud) products.Repository {
	return &productRepository{
		db:    db,
		cloud: cloud,
	}
}

func (repo *productRepository) Create(ctx context.Context, data products.Product) error {
	var inputDB = new(Product)
	inputDB.Name = data.Name
	inputDB.Price = data.Price
	inputDB.Stock = data.Stock
	inputDB.Condition = data.Condition
	inputDB.Description = data.Description
	inputDB.UserId = data.UserId

	for i := 0; i < len(data.Images); i++ {
		url, err := repo.cloud.Upload(ctx, "products", data.Images[i].ImageRaw)
		if err != nil {
			return err
		}

		image := Image{
			ImageURL: *url,
		}

		switch i {
		case 0:
			inputDB.ThumbnailUrl = image.ImageURL
		}

		inputDB.Images = append(inputDB.Images, image)
	}

	if err := repo.db.Create(inputDB).Error; err != nil {
		return err
	}

	return nil
}
