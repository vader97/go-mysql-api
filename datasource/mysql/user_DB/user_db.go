package userdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_username = "MYSQL_USERNAME"
	mysql_passward = "MYSQL_PASSWORD"
	mysql_host     = "MYSQL_HOST"
	mysql_schema   = "MYSQL_SCHEMA"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_username)
	password = os.Getenv(mysql_passward)
	host     = os.Getenv(mysql_host)
	schema   = os.Getenv(mysql_schema)
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "bookmyshow", "127.0.0.1:3306", "go-mysql-api",
	)
	fmt.Println(dataSourceName)
	// var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database Connected Successfully")
}
