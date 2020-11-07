package users

import (
	"fmt"

	"github.com/hikaru-sh/bookstore_users_api/utils/errors"
)
//dao => data access object 
var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DataCreated = result.DataCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
  if current != nil {
		if current.Email != user.Email {
			return errors.NewNotFoundError(fmt.Sprintf("email %s already registered.", user.Email))
		}
    return errors.NewNotFoundError(fmt.Sprintf("user %d already exisits", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}