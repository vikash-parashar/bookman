package utils

import (
	"bookman/model"
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(user model.User) error {
	if user.FullName == "" {
		return errors.New("first name and last name is required")
	}
	if user.Email == "" {
		return errors.New("author name is required")
	}
	if err := ValidatePassword(user.Password); err != nil {
		return err
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
	if !strings.Contains(email, "@") {
		return "", errors.New("invalid email address")
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", errors.New("invalid email address format")
	}

	if parts[1] != "gmail.com" {
		return "", errors.New("provided email address is not a Gmail address")
	}

	return parts[0], nil
}
func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("invalid email address")
	} else if email == "" {
		return errors.New("invalid email address")
	}
	return nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CompaireHash(hashPass string, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainText))
	return err == nil
}
func ValidatePassword(password string) error {
	// Check the password length
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// Regular expression to check for at least one symbol
	symbolRegex := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\}\\|;:'",<>\./?]`)
	if !symbolRegex.MatchString(password) {
		return errors.New("password must contain at least one symbol")
	}

	return nil
}
