package controls

import (
	"bookman/database"
	"bookman/model"
	"bookman/utils"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("book_man ")

func InsertUser(db *sql.DB, user model.User) (res string, err error) {
	// create Table
	_, err = db.Exec(database.UserTable)
	if err != nil {
		return "", err
	}
	// Inser User in DB
	var userId int
	userName, err := utils.ExtractUsernameFromEmail(user.Email)
	if err != nil {
		return "", err
	}
	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	err = db.QueryRow(database.InsertUserIn, user.FullName, userName, user.Email, hashPass, user.MobileNo, user.Role, database.CurentTime).Scan(&userId)
	if err != nil {
		return "", errors.New("email Id and mobile number is already exits")
	}
	res = fmt.Sprintf("Your User ID is %d with provided email %v.", userId, user.Email)
	return res, nil
}
func GernateJwt(w http.ResponseWriter, db *sql.DB, payload model.User) (string, error) {
	var user model.User
	err := db.QueryRow(database.LogerDetail, payload.Email).Scan(&user.UserID, &user.FullName, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return "", errors.New("email id dose not exits with us")
	}
	validate := utils.CompaireHash((user.Password), (payload.Password))
	if !validate {
		return "", errors.New("invalid passsword")
	}
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := model.Credentials{
		Email: payload.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    signedToken,
		Expires:  expirationTime,
		HttpOnly: true, 
		Secure:   true, 
	})
	return signedToken, nil
}
