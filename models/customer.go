// models/customer.go
package models

type Customer struct {
    ID              uint           `gorm:"primaryKey" json:"id"`
    Name            string         `gorm:"type:varchar(255);not null" json:"name"`
    Email           string         `gorm:"type:varchar(255);unique;not null" json:"email"`
    Address         string         `gorm:"type:text" json:"address"`
    // Add more fields as needed
}
