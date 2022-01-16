package user

import (
	"go-self/go-mysql-api/models/user"
	"go-self/go-mysql-api/services"
	"go-self/go-mysql-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO hadle error for Unmarshling the JsonBody
		restErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//fmt.Println(saveErr)
	c.JSON(http.StatusOK, *result)
}

func Get(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		//fmt.Println(err)
		userErr := errors.NewBadRequestErr("invalid userID")
		c.JSON(userErr.Status, userErr)
		return
	}
	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {
	var user user.User
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		//fmt.Println(err)
		userErr := errors.NewBadRequestErr("invalid userID")
		c.JSON(userErr.Status, userErr)
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		resterr := errors.NewBadRequestErr("invalid json body")
		c.JSON(resterr.Status, resterr)
	}
	isPartial := c.Request.Method == http.MethodPatch
	user.Id = userID
	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

func Delete(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		uidErr := errors.NewBadRequestErr("invalid user_id")
		c.JSON(uidErr.Status, uidErr)
		return
	}
	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
	return
}
