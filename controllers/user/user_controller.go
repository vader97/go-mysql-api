package user

import (
	"go-self/go-mysql-api/models/user"
	"go-self/go-mysql-api/services"
	"go-self/go-mysql-api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO hadle error for Unmarshling the JsonBody
		restErr := errors.NewBadRequestErr("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(&user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//fmt.Println(saveErr)
	c.JSON(http.StatusOK, *result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusOK, "GETUSERß")
}
