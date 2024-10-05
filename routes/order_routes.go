package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mdtosif/lumel/controllers"
)

func OrderRoutes(r *gin.Engine) {
  
    r.POST("/orders/import", controllers.AddOrdersFromCSV)
}
