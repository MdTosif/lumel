package main

import (
	"github.com/mdtosif/lumel/config"
	"github.com/mdtosif/lumel/models"
	"github.com/mdtosif/lumel/routes"
)

func main() {
    // Initialize the database connection
    config.InitDatabase()

    // Auto-migrate the Order model to create the table if it doesn't exist
    config.DB.AutoMigrate(&models.Order{})

    // Initialize the router and start the server
    r := routes.InitRoutes()

    // Start the server on port 8080
    r.Run(":8080")  // You can change the port if needed
}
