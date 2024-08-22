package models

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID          uuid.UUID `gorm:"type:uuid;not null"`
	ShowtimeID      uuid.UUID `gorm:"type:uuid;not null"`
	SeatNumbers     string    `gorm:"type:text"` // Changed from []int to string
	ReservationTime time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
