package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
    r := gin.Default()

    // Register order-related routes
    OrderRoutes(r)

    return r
}
