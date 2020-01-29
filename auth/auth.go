package auth

import (
	"github.com/aleale2121/go-webapp/models"
	"github.com/aleale2121/go-webapp/utils"
	"errors"
	"strings"
)

var (
	ErrEmailNotFound = errors.New("Email doenot exitst")
	ErrInvalidPassword = errors.New("Invalid password")
	ErrEmptyFields = errors.New("Fields cannot be empty")
)

func Singin(email, password string) (models.User, error) {
	err := validateFields(strings.ToLower(email), password)
	if err != nil {
		return models.User{}, err
	}
	user, err := models.GetUserByEmail(email)
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