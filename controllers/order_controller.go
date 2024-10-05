package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mdtosif/lumel/config"
	"github.com/mdtosif/lumel/models"
)

// Get all orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	result := config.DB.Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Get order by ID
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	// Convert the id from string to uint
	orderID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	result := config.DB.First(&order, orderID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Get order by ID
func AddOrdersFromCSVfunc(c *gin.Context) {
    // Accept a file from the form
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
        return
    }

    // Open the uploaded file
    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the file"})
        return
    }
    defer src.Close()

    // Create a new CSV reader
    reader := csv.NewReader(src)

    // Read all records from the CSV
    records, err := reader.ReadAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading CSV data"})
        return
    }

    // Process the CSV records (for demonstration, we're just printing them)
    for _, record := range records {
        fmt.Println(record)
    }

    // Respond with a success message
    c.JSON(http.StatusOK, gin.H{"message": "CSV file processed successfully"})
}