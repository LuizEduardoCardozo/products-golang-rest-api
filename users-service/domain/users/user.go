package users

import (
	"fmt"
	"strings"

	usersdb "github.com/LuizEduardoCardozo/catalog-api/users-service/datasources/mysql/users_db"
	"github.com/LuizEduardoCardozo/catalog-api/users-service/utils/date_utils"
	"github.com/LuizEduardoCardozo/catalog-api/users-service/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}

var (
	mockedUsersDatabase = make(map[int64]*User)
)

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("email field cannot to be empty")
	}
	return nil
}

func (user *User) Get() *errors.RestError {
	if err := usersdb.UsersDB.Ping(); err != nil {
		panic(err)
	}

	foundUser := mockedUsersDatabase[user.Id]
	if foundUser == nil {
		return errors.NewNoContentError(fmt.Sprintf("user %d not found", user.Id), "user_not_found")
	}
	user.Id = foundUser.Id
	user.Email = foundUser.Email
	user.LastName = foundUser.LastName
	user.FirstName = foundUser.FirstName
	user.DateCreated = foundUser.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := usersdb.UsersDB.Prepare(usersdb.QInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	user.Id = userId

	return nil
}
