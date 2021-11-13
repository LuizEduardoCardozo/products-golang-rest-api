package services

import (
	"github.com/LuizEduardoCardozo/catalog-api/users-service/domain/users"
	"github.com/LuizEduardoCardozo/catalog-api/users-service/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if validationError := user.Validate(); validationError != nil {
		return nil, validationError
	}
	if saveError := user.Save(); saveError != nil {
		return nil, saveError
	}
	return &user, nil
}

func GetAllUsers() (*[]users.User, *errors.RestError) {
	users, err := users.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int64) (*users.User, *errors.RestError) {
	user := users.User{Id: id}
	if getError := user.Get(); getError != nil {
		return nil, getError
	}
	return &user, nil
}
