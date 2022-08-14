package models

import (
	"errors"
	"rscm/src/security"
	"strings"
)

type User struct {
	User_id  uint64 `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (user *User) Prepare() error {
	if error := user.validateFields(); error != nil {
		return error
	}

	user.formatFields()

	return nil
}

func (user *User) validateFields() error {
	if user.Username == "" {
		return errors.New("Username can't be empty")
	}

	if user.Name == "" {
		return errors.New("Name can't be empty")
	}

	if user.Email == "" {
		return errors.New("Email can't be empty")
	}

	if user.Password == "" {
		return errors.New("Password can't be empty")
	}

	return nil
}

func (user *User) formatFields() {
	user.Username = strings.TrimSpace(user.Username)
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = user.HashPassword(user.Password)

}

func (user *User) HashPassword(password string) (hashedPassword string) {
	hashPassword, error := security.HashPassword(user.Password)
	if error != nil {
		return string(error.Error())
	}
	hashedPassword = string(hashPassword)

	return hashedPassword

}
