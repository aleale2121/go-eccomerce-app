package auth

import (
	"github.com/aleale2121/go-eccomerce-app/models"
	"github.com/aleale2121/go-eccomerce-app/utils"
	"errors"
	"strings"
)

var (
	ErrEmailNotFound = errors.New("Email doenot exitst")
	ErrInvalidPassword = errors.New("Invalid password")
	ErrEmptyFields = errors.New("Fields cannot be empty")
)

func  Singin(s models.Store,email, password string) (models.User, error) {
	err := validateFields(strings.ToLower(email), password)
	if err != nil {
		return models.User{}, err
	}
	user, err := s.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, ErrEmailNotFound
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}
	return user, nil
}

func validateFields(email, password string) error {
	if models.IsEmpty(models.Trim(email)) || models.IsEmpty(password) {
		return ErrEmptyFields
	}
	if !models.IsEmail(email) {
		return models.ErrInvalidEmail
	}
	return nil
}