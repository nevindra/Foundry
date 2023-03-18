package model

import (
	"gorm.io/gorm"
	"time"
)

// User model for user
// This model will have ID, Email, Password, Verified, created_at
// ID is the primary key and Created_at is the timestamp of when the user is created
type User struct {
	gorm.Model
	ID        string    `json:"id"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Datasets  []Dataset `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Projects  []Project `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
