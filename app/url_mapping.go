package app

import (
	"go-self/go-mysql-api/controllers/ping"
	"go-self/go-mysql-api/controllers/user"
)

func MapUrl() {
	router.GET("/ping", ping.Ping)
	router.GET("/user/:userID", user.GetUser)
	router.POST("/createuser", user.CreateUser)
}
