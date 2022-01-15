package user

import (
	"fmt"
	"go-self/go-mysql-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.Lastname = result.Lastname
	user.Email = result.Email
	return nil
}

func (user *User) Save() *errors.RestErr {
	if userDB[user.Id] != nil {
		return errors.NewBadRequestErr("user already exist")
	}
	userDB[user.Id] = user
	return nil
}
