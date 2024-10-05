// models/payment.go
package models

type Payment struct {
    ID              uint   `gorm:"primaryKey" json:"id"`
    Method          string `gorm:"type:varchar(50)" json:"method"`
}
