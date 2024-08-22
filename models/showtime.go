package models

import (
	"time"

	"github.com/google/uuid"
)

type Showtime struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	MovieID        uuid.UUID `gorm:"not null;" json:"movieId"`
	StartTime      time.Time `gorm:"not null;" json:"startTime"`
	EndTime        time.Time `gorm:"not null;" json:"endTime"`
	AvailableSeats int       `gorm:"not null;" json:"availableSeats"`
	Price          float64   `gorm:"not null;" json:"price"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
