package user

import (
	"fmt"
	userdb "go-self/go-mysql-api/datasource/mysql/user_DB"
	"go-self/go-mysql-api/utils/errors"
	"strings"
	"time"
)

const (
	insertUserQuery = "INSERT INTO users(first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
	getUserQuery    = "SELECT id,first_name,last_name,email,created_at from users WHERE id = ?;"
	updateUserQuery = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	deleteUserQuery = "DELETE FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(getUserQuery)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.Lastname, &user.Email, &user.CreatedAt); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return errors.NewNotFoundErr("no user exist for given id")
		}
		return errors.NewInternalServerErr(err.Error())
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(insertUserQuery)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()
	user.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")

	insertResult, err := stmt.Exec(user.FirstName, user.Lastname, user.Email, user.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestErr("email is already registered")
		}
		return errors.NewInternalServerErr(
			fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerErr(
			fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}
	user.Id = userID
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(updateUserQuery)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.Lastname, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerErr("error occured while updating the usr")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(deleteUserQuery)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		return errors.NewInternalServerErr("error occured while deleting the user ")
	}
	return nil
}
