package controls

import (
	"bookman/database"
	"bookman/model"
	"database/sql"
	"time"
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
	err = db.QueryRow(database.InsertUserIn, user.FullName, user.Username, user.Email, user.MobileNo, user.Role, time.Now().Format("15:04:05 Monday 01-02-2006")).Scan(&userId)
	if err != nil {
		return res, err
	}
	return model.User{
		UserID:    userId,
		FullName:  user.FullName,
		Email:     user.Email,
		MobileNo:  user.MobileNo,
		Role:      user.Role,
		CreatedAt: database.CurentTime,
	}, nil
}
