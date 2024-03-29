package models

import (
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"strings"
)

var (
	ErrRequiredFirstName = errors.New("Name required")
	ErrRequiredLastName  = errors.New("FatherName required")
	ErrRequiredEmail     = errors.New("Email required")
	ErrInvalidEmail      = errors.New("Email invalid")
	ErrRequiredPassword  = errors.New("Password required")
	ErrMaxLimit          = errors.New("Maximum limit of caracteres reached")
	ErrEmailTaken        = errors.New("Email already taken")
)

func IsEmpty(attr string) bool {
	return attr == "" 
}

func Trim(attr string) string {
	return strings.TrimSpace(attr)
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	return err == nil
}

func Max(attr string, lim int) bool {
	return len(attr) <= lim 
}

func ValidateLimitFields(user User) (User, error) {
	if !Max(user.FirstName, 15) || !Max(user.LastName, 20) || !Max(user.Email, 40) || !Max(user.Password, 100) {
		return User{}, ErrMaxLimit
	}
	return user, nil
}

func (s Store) VerifyEmail(email string) (bool, error) {

	sql := "select count(email) from users where email = $1"
	rs, err := s.DB.Query(sql, email)
	if err != nil {
		return false, err
	}
	defer rs.Close()
	var count int64
	if rs.Next() {
		err := rs.Scan(&count)
		if err != nil {
			return false, err
		}
	}
	if count > 0 {
		return false, ErrEmailTaken
	}
	return true, nil
}

func (s Store)  ValidateNewUser(user User) (User, error) {
	user, err := ValidateLimitFields(user)
	if err != nil {
		return user, err
	}
	user.FirstName = Trim(user.FirstName)
	user.LastName = Trim(user.LastName)
	user.Email = Trim(strings.ToLower(user.Email))
	if IsEmpty(user.FirstName) {
		return User{}, ErrRequiredFirstName
	}
	if IsEmpty(user.LastName) {
		return User{}, ErrRequiredLastName
	}
	if IsEmpty(user.Email) {
		return User{}, ErrRequiredEmail
	}
	if !IsEmail(user.Email) {
		return User{}, ErrInvalidEmail
	}
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}
	_, err = s.VerifyEmail(user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s Store) Count(table string) (int64, error) {

	sql := fmt.Sprintf("select count(*) from %s", table)
	var count int64
	err := s.DB.QueryRow(sql).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
