package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	QInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	UsersDB *sql.DB
)

func init() {

	connUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "localhost", "users")

	var err error
	UsersDB, err = sql.Open("mysql", connUrl)
	if err != nil {
		panic(err)
	}

	if err := UsersDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database successfully connected")
}
