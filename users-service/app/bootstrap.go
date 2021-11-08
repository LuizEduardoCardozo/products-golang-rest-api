package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Bootstrap(port int) {
	RouteMap(router)
	router.Run(fmt.Sprintf(":%d", port))
}
