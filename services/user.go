package services

import (
	"go-self/go-mysql-api/models/user"
	"go-self/go-mysql-api/utils/errors"
)

func CreateUser(user user.User) (*user.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*user.User, *errors.RestErr) {
	var user user.User
	user.Id = userID
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}
