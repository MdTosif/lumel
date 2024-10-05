package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	for _,data := range csvData{
		fmt.Println("orderId:", data.OrderID)
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "CSV file processed successfully"})
}
