package app

import (
	"github.com/LuizEduardoCardozo/catalog-api/users-service/controllers"
	"github.com/gin-gonic/gin"
)

func RouteMap(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:userId", controllers.GetUser)
}
