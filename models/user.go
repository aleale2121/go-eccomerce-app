package models

import (
	"github.com/aleale2121/go-eccomerce-app/utils"
)

type User struct {
	Id        uint64
	FirstName string
	LastName  string
	Email     string
	Password  string
	Status    string
}

func  (s Store) NewUser(user User) (bool, error) {
	user, err := s.ValidateNewUser(user)
	if err != nil {
		return false, err
	}

	sql := "insert into users (firstname, lastname, email, password) values ($1, $2, $3, $4)"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	hash, err := utils.Hash(user.Password)
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, hash)
	if err != nil {
		return false, err
	}
	return true, err
}

func  (s Store) GetUserByEmail(email string) (User, error) {

	sql := "select * from users where email = $1"
	rs, err := s.DB.Query(sql, email)
	if err != nil {
		return User{}, err
	}
	defer rs.Close()
	var user User
	if rs.Next() {
		err := rs.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func  (s Store) GetUsers() ([]User, error) {

	sql := "select * from users"
	rs, err := s.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var users []User
	for rs.Next() {
		var user User
		err := rs.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
