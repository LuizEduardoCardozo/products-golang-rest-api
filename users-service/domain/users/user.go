package users

import (
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

func GetAllUsers() (*[]User, *errors.RestError) {
	var foundUsers []User

	stmt, err := usersdb.UsersDB.Prepare(usersdb.QGetAllUsers)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	defer stmt.Close()

	queryResult, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	for queryResult.Next() {
		var user User
		if err := queryResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		foundUsers = append(foundUsers, user)
	}

	return &foundUsers, nil
}

func (user *User) Get() *errors.RestError {
	stmt, err := usersdb.UsersDB.Prepare(usersdb.QGetUserById)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	defer stmt.Close()

	query, err := stmt.Query(user.Id)
	for query.Next() {
		if err := query.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
			return errors.NewBadRequestError(err.Error())
		}
	}

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
