package service_test

import (
	"bilo/features/users"
	"bilo/features/users/mocks"
	"bilo/features/users/service"
	encMock "bilo/helper/encrypt/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceRegister(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var enc = encMock.NewHashInterface(t)
	var srv = service.New(repo, enc)

	t.Run("invalid name", func(t *testing.T) {
		var caseData = users.User{
			Name:     "",
			Username: "galihpryg",
			Email:    "galih@mail.com",
			Password: "test",
			Hp:       "081229081229",
		}

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "name")
	})

	t.Run("invalid phone", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Username: "galihpryg",
			Email:    "galihp83@gmail.com",
			Password: "test",
			Hp:       "",
		}

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "phone")
	})

	t.Run("invalid email", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Username: "galihpryg",
			Email:    "",
			Password: "test",
			Hp:       "081229081229",
		}

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "email")
	})

	t.Run("invalid password", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Username: "galihpryg",
			Email:    "galih@gmail.com",
			Password: "",
			Hp:       "081229081229",
		}

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "password")
	})

	t.Run("error from encrypt", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Username: "galihpryg",
			Email:    "galih@gmail.com",
			Password: "test",
			Hp:       "081229081229",
		}

		enc.On("HashPassword", caseData.Password).Return("", errors.New("HASH - something wrong when hashing password")).Once()

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "HASH - something wrong when hashing password")

		enc.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Username: "galihpryg",
			Email:    "galih@gmail.com",
			Password: "test",
			Hp:       "081229081229",
			Image:    "",
		}

		enc.On("HashPassword", caseData.Password).Return("secret", nil).Once()

		caseData.Password = "secret"
		repo.On("Register", caseData).Return(errors.New("some error from repository")).Once()

		caseData.Password = "test"
		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "some error from repository")

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Username: "galihpryg",
			Email:    "galih@gmail.com",
			Password: "test",
			Hp:       "081229081229",
			Image:    "",
		}

		enc.On("HashPassword", caseData.Password).Return("secret", nil).Once()

		caseData.Password = "secret"
		repo.On("Register", caseData).Return(nil).Once()

		caseData.Password = "test"
		err := srv.Register(caseData)

		assert.NoError(t, err)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})
}
