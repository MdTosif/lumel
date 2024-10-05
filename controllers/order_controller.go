package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdtosif/lumel/config"
	"github.com/mdtosif/lumel/service"
)

// Get order by ID
func AddOrdersFromCSV(c *gin.Context) {
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

	csvData, err := service.ReadCSV(src)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the file"})
		return
	}

	for _, data := range csvData {
		fmt.Println("orderId:", data.OrderID)
		err := service.ImportOrder(config.DB, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the file"})
			return
		}
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "CSV file processed successfully"})
}

type TopProductsRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Limit     int       `json:"limit" binding:"required"`
}

// TopNProductsOverall handles the request for overall top N products
func TopNProductsOverall(c *gin.Context) {
	var request TopProductsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var products []struct {
		ProductID    string `json:"product_id"`
		QuantitySold int    `json:"quantity_sold"`
	}
	config.DB.Debug()
	// Query to get top N products overall based on quantity sold
	if err := config.DB.Table("orders").
		Select("product_id, SUM(quantity_sold) AS quantity_sold").
		Where("date_of_sale BETWEEN ? AND ?", request.StartDate, request.EndDate).
		Group("product_id").
		Order("quantity_sold DESC").
		Limit(request.Limit).
		Scan(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// TopNProductsByCategory handles the request for top N products by category
func TopNProductsByCategory(c *gin.Context) {
	var request TopProductsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := c.Param("category")

	var products []struct {
		ProductID    string `json:"product_id"`
		QuantitySold int    `json:"quantity_sold"`
	}

	// Query to get top N products by category based on quantity sold
	if err := config.DB.Table("orders").
		Select("product_id, SUM(quantity_sold) AS quantity_sold").
		Joins("JOIN products ON products.id = orders.product_id").
		Where("products.category iLike ? AND date_of_sale BETWEEN ? AND ?", "%" +category+ "%", request.StartDate, request.EndDate).
		Group("product_id").
		Order("quantity_sold DESC").
		Limit(request.Limit).
		Scan(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// TopNProductsByRegion handles the request for top N products by region
func TopNProductsByRegion(c *gin.Context) {
	var request TopProductsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	region := c.Param("region")

	var products []struct {
		ProductID    string `json:"product_id"`
		QuantitySold int    `json:"quantity_sold"`
	}

	// Query to get top N products by region based on quantity sold
	if err := config.DB.Table("orders").
		Select("product_id, SUM(quantity_sold) AS quantity_sold").
		Joins("JOIN customers ON customers.id = orders.customer_id").
		Where("customers.address iLike ? AND date_of_sale BETWEEN ? AND ?", "%"+region+"%", request.StartDate, request.EndDate).
		Group("product_id").
		Order("quantity_sold DESC").
		Limit(request.Limit).
		Scan(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
