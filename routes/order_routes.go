package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mdtosif/lumel/controllers"
)

func OrderRoutes(r *gin.Engine) {
  
    r.POST("/orders/import", controllers.AddOrdersFromCSV)
    // Routes for top N products
    r.POST("/orders/top-overall", controllers.TopNProductsOverall)
    r.POST("/orders/top-by-category/:category", controllers.TopNProductsByCategory)
    r.POST("/orders/top-by-region/:region", controllers.TopNProductsByRegion)

}
