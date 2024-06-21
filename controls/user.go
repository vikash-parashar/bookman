package controls

import (
	"bookman/database"
	"bookman/model"
	"bookman/utils"
	"database/sql"
	"errors"
)

func InsertUser(db *sql.DB, user model.User) (res model.User, err error) {
	// create Table
	_, err = db.Exec(database.UserTable)
	if err != nil {
		// err = errors.New("failed to create user table")
		return res, err
	}

	// Insert User
	// err = IsUSerRegistered(db, user.BookName)
	// if err != nil {

	// 	return response, err
	// }
	var userId int
	userName, err := utils.ExtractUsernameFromEmail(user.Email)
	if err != nil {
		return res, err
	}

	err = db.QueryRow(database.InsertUserIn, user.FullName, userName, user.Email, user.MobileNo, user.Role, database.CurentTime).Scan(&userId)
	if err != nil {
		return res, errors.New("email Id and mobile number is already exits")
	}
	return model.User{
		UserID:    userId,
		FullName:  user.FullName,
		Username:  userName,
		Email:     user.Email,
		MobileNo:  user.MobileNo,
		Role:      user.Role,
		CreatedAt: database.CurentTime,
	}, nil
}
