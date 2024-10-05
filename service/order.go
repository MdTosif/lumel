package service

import (
	"time"

	"github.com/mdtosif/lumel/models"
	"gorm.io/gorm"
)

// var DB *gorm.DB = config.DB

func ImportOrder(DB *gorm.DB, order OrderCSV) error {
	// Map the OrderCSV to the Product model
	product := models.Product{
		ID:       order.ProductID,
		Name:     order.ProductName,
		Category: order.Category,
		Price:    order.UnitPrice,
	}

	// Check if product exists and create or update accordingly
	if err := DB.Where("id = ?", product.ID).First(&product).Error; err != nil {
		// Create the product if not found
		if err == gorm.ErrRecordNotFound {
			if err := DB.Create(&product).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Map the OrderCSV to the Customer model
	customer := models.Customer{
		ID:      order.CustomerID,
		Name:    order.CustomerName,
		Email:   order.CustomerEmail,
		Address: order.CustomerAddress,
	}

	// Check if customer exists and create or update accordingly
	if err := DB.Where("id = ?", customer.ID).First(&customer).Error; err != nil {
		// Create the customer if not found
		if err == gorm.ErrRecordNotFound {
			if err := DB.Create(&customer).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Map the OrderCSV to the Payment model
	payment := models.Payment{
		Method: order.PaymentMethod,
	}

	// Check if payment method exists and create or update accordingly
	if err := DB.Where("method = ?", payment.Method).First(&payment).Error; err != nil {
		// Create the payment method if not found
		if err == gorm.ErrRecordNotFound {
			if err := DB.Create(&payment).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Create the Order
	orderRecord := models.Order{
		ID:           order.OrderID,
		ProductID:    product.ID,
		CustomerID:   customer.ID,
		PaymentID:    payment.ID,
		QuantitySold: order.QuantitySold,
		Discount:     order.Discount,
		ShippingCost: order.ShippingCost,
		DateOfSale:   parseDate(order.DateOfSale),
	}

	// Check if payment method exists and create or update accordingly
	if err := DB.Where("id = ?", orderRecord.ID).First(&orderRecord).Error; err != nil {
		// Create the payment method if not found
		if err == gorm.ErrRecordNotFound {
			if err := DB.Create(&orderRecord).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func parseDate(dateStr string) time.Time {
	// Here you can parse the date string from the CSV to time.Time format
	parsedTime, _ := time.Parse("2006-01-02", dateStr) // Adjust format as needed
	return parsedTime
}
