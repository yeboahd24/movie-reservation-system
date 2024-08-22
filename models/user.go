package models

import (
	"time"

	"github.com/google/uuid"
)

type Role int

const (
	AdminRole Role = iota
	RegularUserRole
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"not null;"`
	Email     string    `gorm:"unique;not null;"`
	Password  string    `gorm:"not null;"`
	Role      Role      `gorm:"not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
