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

func TestUserServiceLogin(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var enc = encMock.NewHashInterface(t)
	var srv = service.New(repo, enc)

	t.Run("invalid email", func(t *testing.T) {
		var caseData = users.User{
			Email:    "",
			Password: "test",
		}

		result, err := srv.Login(caseData.Email, caseData.Password)

		assert.ErrorContains(t, err, "email")
		assert.Nil(t, result)
	})

	t.Run("invalid password", func(t *testing.T) {
		var caseData = users.User{
			Email:    "galih@gmail.com",
			Password: "",
		}

		result, err := srv.Login(caseData.Email, caseData.Password)

		assert.ErrorContains(t, err, "password")
		assert.Nil(t, result)
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = users.User{
			Email:    "galih@gmail.com",
			Password: "test",
		}

		repo.On("Login", caseData.Email).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Login(caseData.Email, caseData.Password)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		var caseData = users.User{
			Email:    "galih@gmail.com",
			Password: "wrongpassword",
		}

		var caseResult = users.User{
			Id:       1,
			Name:     "Galih",
			Username: "galihpryg",
			Email:    "galih@gmail.com",
		}

		repo.On("Login", caseData.Email).Return(&caseResult, nil).Once()
		enc.On("Compare", caseResult.Password, caseData.Password).Return(errors.New("wrong password")).Once()
		res, err := srv.Login(caseData.Email, caseData.Password)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)

		assert.ErrorContains(t, err, "wrong password")
		assert.Nil(t, res)

	})

	t.Run("success", func(t *testing.T) {
		var caseData = users.User{
			Email:    "galih@gmail.com",
			Password: "test",
		}

		var caseResult = users.User{
			Id:       1,
			Name:     "Galih",
			Username: "galihpryg",
			Email:    "galih@gmail.com",
		}

		repo.On("Login", caseData.Email).Return(&caseResult, nil).Once()
		enc.On("Compare", caseResult.Password, caseData.Password).Return(nil).Once()
		res, err := srv.Login(caseData.Email, caseData.Password)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, "Galih", res.Name)
	})
}
