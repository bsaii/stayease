package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomNumber  string    `gorm:"not null" json:"room_number"`             // A unique identifier for the room (e.g., room number or code)
	Type        string    `gorm:"not null" json:"type"`                    // Type of room (e.g., Single, Double, Suite)
	Description string    `gorm:"not null" json:"description"`             // Brief description of the room (features, amenities, etc.)
	Capacity    int       `gorm:"not null" json:"capacity"`                // Maximum number of occupants the room can accommodate
	Price       float32   `gorm:"default:0.00;not null" json:"price"`      // Price per night for booking the room
	IsBooked    bool      `gorm:"default:false;not null" json:"is_booked"` // Indicates whether the room is currently booked
	BookedDates []Booking `json:"booked_dates"`
}
