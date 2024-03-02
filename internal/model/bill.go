package model

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	CustomerName string  `gorm:"not null" json:"customer_name"`       // Name of the customer associated with the bill
	Amount       float64 `gorm:"default:0.00;not null" json:"amount"` // Amount of the bill
	Description  string  `gorm:"not null" json:"description"`         // Description or additional information about the bill
	Paid         bool    `gorm:"default:false;not null" json:"paid"`  // Indicates whether the bill has been paid, default is false
}
