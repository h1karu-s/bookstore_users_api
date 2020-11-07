package users

import (
	"strings"

	"github.com/hikaru-sh/bookstore_users_api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
 	LastName    string `json:"last_name"`
	Email       string `josn:"email"`
	DataCreated string `json:"data_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return  errors.NewBadRequestError("invalid email address.")
	}
	return nil
}