package service

import (
	"bilo/features/users"
	"bilo/helper/encrypt"
	"errors"
)

type userService struct {
	repo users.Repository
	hash encrypt.HashInterface
}

func New(repo users.Repository, hash encrypt.HashInterface) users.Service {
	return &userService{
		repo: repo,
		hash: hash,
	}
}

func (service *userService) Register(newUser users.User) error {
	if newUser.Name == "" {
		return errors.New("validate: name can't be empty")
	}
	if newUser.Hp == "" {
		return errors.New("validate: phone can't be empty")
	}
	if newUser.Email == "" {
		return errors.New("validate: email can't be empty")
	}
	if newUser.Password == "" {
		return errors.New("validate: password can't be empty")
	}

	encrypt, err := service.hash.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = encrypt

	if err := service.repo.Register(newUser); err != nil {
		return err
	}

	return nil
}

func (us *userService) Login(email string, password string) (users.User, error) {
	panic("on progress")
}
