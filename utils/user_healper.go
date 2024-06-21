package utils

import (
	"bookman/model"
	"errors"
	"strings"
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
func ExtractUsernameFromEmail(email string) (string, error) {
	// Check if the email address contains "@"
	if !strings.Contains(email, "@") {
		return "", errors.New("invalid email address")
	}

	// Split the email address into username and domain
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", errors.New("invalid email address format")
	}

	// Check if the domain is "gmail.com"
	if parts[1] != "gmail.com" {
		return "", errors.New("provided email address is not a Gmail address")
	}

	return parts[0], nil
}
