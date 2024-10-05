package service

import (
	"time"

	"github.com/mdtosif/lumel/models"
	"gorm.io/gorm"
)

func ImportOrder(db *gorm.DB, order OrderCSV) error {
	// Map the OrderCSV to the Product model
	product := models.Product{
		Name:     order.ProductName,
		Category: order.Category,
		Price:    order.UnitPrice,
	}

	// Check if product exists and create or update accordingly
	if err := db.Where("name = ?", product.Name).First(&product).Error; err != nil {
		// Create the product if not found
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&product).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Map the OrderCSV to the Customer model
	customer := models.Customer{
		Name:    order.CustomerName,
		Email:   order.CustomerEmail,
		Address: order.CustomerAddress,
	}

	// Check if customer exists and create or update accordingly
	if err := db.Where("email = ?", customer.Email).First(&customer).Error; err != nil {
		// Create the customer if not found
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&customer).Error; err != nil {
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
	if err := db.Where("method = ?", payment.Method).First(&payment).Error; err != nil {
		// Create the payment method if not found
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&payment).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Create the Order
	orderRecord := models.Order{
		ProductID:    product.ID,
		CustomerID:   customer.ID,
		PaymentID:    payment.ID,
		QuantitySold: order.QuantitySold,
		Discount:     order.Discount,
		ShippingCost: order.ShippingCost,
		DateOfSale:   parseDate(order.DateOfSale),
	}

	return db.Create(&orderRecord).Error
}

func parseDate(dateStr string) time.Time {
	// Here you can parse the date string from the CSV to time.Time format
	parsedTime, _ := time.Parse("2006-01-02", dateStr) // Adjust format as needed
	return parsedTime
}
