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

func UpdateUser(partial bool, user user.User) (*user.User, *errors.RestErr) {
	existUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if partial {
		if user.FirstName != "" {
			existUser.FirstName = user.FirstName
		}
		if user.Lastname != "" {
			existUser.Lastname = user.Lastname
		}
		if user.Email != "" {
			existUser.Email = user.Email
		}

	} else {
		existUser.FirstName = user.FirstName
		existUser.Lastname = user.Lastname
		existUser.Email = user.Email
	}
	if err := existUser.Update(); err != nil {
		return nil, err
	}
	return existUser, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := user.User{Id: userId}
	err := user.Delete()
	if err != nil {
		return err
	}
	return nil
}
