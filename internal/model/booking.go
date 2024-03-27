package model

import (
	"time"
)

type Booking struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	RoomID       uint      `json:"room_id"`                                 // ID of the room being booked                                 // ID of the user making the booking
	CheckInDate  time.Time `json:"check_in_date"`                           // Date and time of check-in
	CheckOutDate time.Time `json:"check_out_date"`                          // Date and time of check-out
	TotalCost    float32   `gorm:"default:0.00;not null" json:"total_cost"` // Total cost of the booking
}
