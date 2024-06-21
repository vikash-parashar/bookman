package utils

import (
	"bookman/model"
	"errors"
)

func ValidateUser(user model.User) error {
	if user.FullName == "" {
		return errors.New("first name and last name is required")
	}
	if user.Email == "" {
		return errors.New("author name is required")
	}
	if user.MobileNo == "" {
		return errors.New("mobile number can not be blank and ten disit")
	}
	if user.Role == "" {
		return errors.New("role is required for register user")
	}
	return nil
}
