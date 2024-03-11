package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	//gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Login     string    `gorm:"unique"`
	Password  string
	IsActive  bool `gorm:"DEFAULT:false"`
	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
