package controllers

import (
	"net/http"
	"strconv"

	"github.com/LuizEduardoCardozo/catalog-api/users-service/domain/users"
	"github.com/LuizEduardoCardozo/catalog-api/users-service/services"
	"github.com/LuizEduardoCardozo/catalog-api/users-service/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		ctx.JSON(restErr.Code, restErr)
		return
	}
	createdUser, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}
	ctx.JSON(http.StatusCreated, createdUser)
	return
}

func GetUser(ctx *gin.Context) {
	userIdParam := ctx.Param("userId")
	userId, parseError := strconv.ParseInt(userIdParam, 10, 64)
	if parseError != nil {
		ctx.JSON(http.StatusBadRequest, errors.NewBadRequestError("userId should be a number"))
		return
	}
	foundUser, err := services.GetUser(userId)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}
	ctx.JSON(http.StatusOK, foundUser)
	return
}
