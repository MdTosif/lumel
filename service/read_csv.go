// service.go
package service

import (
	"io"

	"github.com/gocarina/gocsv"
)

type OrderCSV struct {
	OrderID         int  `csv:"Order ID"`
	ProductID       string  `csv:"Product ID"`
	CustomerID      string  `csv:"Customer ID"`
	ProductName     string  `csv:"Product Name"`
	Category        string  `csv:"Category"`
	Region          string  `csv:"Region"`
	DateOfSale      string  `csv:"Date of Sale"`
	QuantitySold    int     `csv:"Quantity Sold"`
	UnitPrice       float64 `csv:"Unit Price"`
	Discount        float64 `csv:"Discount"`
	ShippingCost    float64 `csv:"Shipping Cost"`
	PaymentMethod   string  `csv:"Payment Method"`
	CustomerName    string  `csv:"Customer Name"`
	CustomerEmail   string  `csv:"Customer Email"`
	CustomerAddress string  `csv:"Customer Address"`
}

func ReadCSV(file io.Reader) ([]OrderCSV, error) {
	var orders []OrderCSV
	

	// Unmarshal the CSV data into the orders slice
	if err := gocsv.Unmarshal(file, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
