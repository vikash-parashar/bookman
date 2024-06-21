package model

import "github.com/golang-jwt/jwt/v4"

type User struct {
	UserID    int    `json:"user_id"`
	FullName  string `json:"full_name"`
	Username  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	MobileNo  string `json:"mobile_no"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type Credentials struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
