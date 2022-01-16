package app

import (
	"go-self/go-mysql-api/controllers/ping"
	"go-self/go-mysql-api/controllers/user"
)

func MapUrl() {
	router.GET("/ping", ping.Ping)
	router.GET("/user/:user_id", user.Get)
	router.POST("/create", user.Create)
	router.PUT("/user/:user_id", user.Update)
	router.PATCH("/user/:user_id", user.Update)
	router.DELETE("user/:user_id", user.Delete)
}
