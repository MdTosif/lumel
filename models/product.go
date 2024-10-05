// models/product.go
package models

type Product struct {
    ID                 string   `gorm:"primaryKey" json:"id"`
    Name               string `gorm:"type:varchar(255);not null" json:"name"`
    Category           string `gorm:"type:varchar(100)" json:"category"`
    Description        string `gorm:"type:text" json:"description"` // Optional
    Price              float64 `json:"price"`                      // Unit price
    // Add more fields as needed
}
