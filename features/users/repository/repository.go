package repository

import (
	"bilo/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Hp       string
	Email    string `gorm:"unique"`
	Password string
	Image    string
	Username string `gorm:"unique"`
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.Repository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) Register(newUser users.User) error {
	var inputDB = new(User)
	inputDB.Name = newUser.Name
	inputDB.Email = newUser.Email
	inputDB.Hp = newUser.Hp
	inputDB.Password = newUser.Password
	inputDB.Username = newUser.Username

	if err := repo.db.Create(inputDB).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) Login(email string) (*users.User, error) {
	var data = new(User)

	if err := repo.db.Where("email = ?", email).First(data).Error; err != nil {
		return nil, err
	}

	var result = new(users.User)
	result.ID = data.ID
	result.Name = data.Name

	return result, nil
}
