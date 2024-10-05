package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
    OrderID         uint           `gorm:"primaryKey" json:"order_id"`
    ProductID       uint           `gorm:"not null" json:"product_id"`
    CustomerID      uint           `gorm:"not null" json:"customer_id"`
    ProductName     string         `gorm:"type:varchar(255);not null" json:"product_name"`
    Category        string         `gorm:"type:varchar(100)" json:"category"`
    Region          string         `gorm:"type:varchar(100)" json:"region"`
    DateOfSale      time.Time      `json:"date_of_sale"`
    QuantitySold    int            `json:"quantity_sold"`
    UnitPrice       float64        `json:"unit_price"`
    Discount        float64        `json:"discount"`
    ShippingCost    float64        `json:"shipping_cost"`
    PaymentMethod   string         `gorm:"type:varchar(50)" json:"payment_method"`
    CustomerName    string         `gorm:"type:varchar(255);not null" json:"customer_name"`
    CustomerEmail   string         `gorm:"type:varchar(255);unique;not null" json:"customer_email"`
    CustomerAddress string         `gorm:"type:text" json:"customer_address"`
    ProductDescription string      `gorm:"type:text" json:"product_description"`  // Optional
    CampaignDetails string         `gorm:"type:text" json:"campaign_details"`      // Optional
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
