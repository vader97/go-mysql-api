package user

import (
	"fmt"
	"go-self/go-mysql-api/utils/errors"
	"time"
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
	user.CreatedAt = result.CreatedAt
	return nil
}

func (user *User) Save() *errors.RestErr {
	if userDB[user.Id] != nil {
		return errors.NewBadRequestErr("user already exist")
	}
	user.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")
	userDB[user.Id] = user
	return nil
}
