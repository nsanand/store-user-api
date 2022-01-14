package users

import (
	"github.com/nsanand/store-user-api/utils/errors"
	"strings"
)

type User struct {
	Id         int64 	`json:"id"`
	FirstName  string 	`json:"first_name"`
	LastName   string	`json:"last_name"`
	Email      string	`json:"email"`
	DateCreate string	`json:"date_create"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("Invalid Email Address")
	}
	return nil
}
