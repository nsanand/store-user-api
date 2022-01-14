package users

import (
	"fmt"
	"github.com/nsanand/store-user-api/datasources/mysql/users_db"
	"github.com/nsanand/store-user-api/utils/date_utils"
	"github.com/nsanand/store-user-api/utils/errors"
	"strings"
)

const (
	emailUnique     = "email_UNIQUE"
	NoRows          = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)


func (user *User) Get() *errors.RestErr {
	//stmt, err := users_db.Client.Prepare(queryGetUser)
	//if err != nil {
	//	return errors.InternalServerError(err.Error())
	//}
	//defer stmt.Close()
	//insertResult := stmt.QueryRow(user.Id)

	result := users_db.Client.QueryRow(queryGetUser, user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreate); err != nil {
		if strings.Contains(err.Error(), NoRows) {
			return errors.NotFoundError(fmt.Sprintf("user id %d not found", user.Id))
		}
		return errors.InternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	user.DateCreate = date_utils.GetNowString()
	//stmt, err := users_db.Client.Prepare(queryInsertUser)
	//if err != nil {
	//	return errors.InternalServerError(err.Error())
	//}
	//defer stmt.Close()
	//insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreate)
	//if err != nil {
	//	return errors.InternalServerError(
	//		fmt.Sprintf("error when trying to save user: %s", err.Error()))
	//}
	insertResult, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreate)
	if err != nil {
		if strings.Contains(err.Error(), emailUnique) {
			return errors.BadRequestError(
				fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.InternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
