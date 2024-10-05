package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mdtosif/lumel/controllers"
)

func OrderRoutes(r *gin.Engine) {
    // Route to get all orders
    r.GET("/orders", controllers.GetOrders)

    // Route to get a specific order by its ID
    r.GET("/orders/:id", controllers.GetOrderByID)
    r.GET("/orders/import", controllers)
}
