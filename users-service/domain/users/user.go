package users

import (
	"fmt"
	"strings"

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
	if mockedUsersDatabase[user.Id] == nil {
		user.DateCreated = date_utils.GetNowString()
		mockedUsersDatabase[user.Id] = user
		return nil
	}
	return errors.NewConflictError("user already exists", "user_already_exists")
}
