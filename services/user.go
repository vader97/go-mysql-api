package services

import (
	"go-self/go-mysql-api/models/user"
	"go-self/go-mysql-api/utils/errors"
)

func CreateUser(user *user.User) (*user.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}
