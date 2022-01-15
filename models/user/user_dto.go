package user

import (
	"go-self/go-mysql-api/utils/errors"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"f_name"`
	Lastname  string `json:"l_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestErr("invalid Email address")
	}
	return nil
}
