// models/order.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
    ID              int           `gorm:"primaryKey" json:"id"`
    ProductID       string           `gorm:"not null" json:"product_id"`
    CustomerID      string           `gorm:"not null" json:"customer_id"`
    PaymentID       uint           `gorm:"not null" json:"payment_id"`
    QuantitySold    int            `json:"quantity_sold"`
    Discount        float64        `json:"discount"`
    ShippingCost    float64        `json:"shipping_cost"`
    DateOfSale      time.Time      `json:"date_of_sale"`
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

    // Relationships
    Product         Product        `gorm:"foreignKey:ProductID" json:"product"`
    Customer        Customer       `gorm:"foreignKey:CustomerID" json:"customer"`
    Payment         Payment        `gorm:"foreignKey:PaymentID" json:"payment"`
}
