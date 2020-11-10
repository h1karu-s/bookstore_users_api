package users

import (
	"fmt"

	"github.com/hikaru-sh/bookstore_users_api/datasources/mysql/users_db"
	"github.com/hikaru-sh/bookstore_users_api/utils/date_utils"
	"github.com/hikaru-sh/bookstore_users_api/utils/errors"
	"github.com/hikaru-sh/bookstore_users_api/utils/mysql_utils"
)

const(
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser   = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

//dao => data access object


func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DataCreated); getErr != nil {
	  return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DataCreated = date_utils.GetNowString()
	insertUser, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DataCreated)
  if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertUser.LastInsertId()
	if err != nil {
    return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s.", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError("error when trying to update user.")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
} 


func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError("error when trying to delete user.")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return  errors.NewInternalServerError("error when trying to delete user.")
	}
	return nil
}


