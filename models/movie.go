package models

import (
	"time"

	"github.com/google/uuid"
)

type Genre string

const (
	ActionGenre         Genre = "Action"
	ComedyGenre         Genre = "Comedy"
	DramaGenre          Genre = "Drama"
	ScienceFictionGenre Genre = "Science Fiction"
)

type Movie struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"not null;"`
	Description string    `gorm:"not null;"`
	PosterImage string    `gorm:"not null;column:poster_image"`
	Genre       Genre     `gorm:"not null;"`
	Duration    int       `json:"duration"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"releaseDate"`
	PosterURL   string    `json:"posterURL"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
