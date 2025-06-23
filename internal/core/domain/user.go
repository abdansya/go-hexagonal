package domain

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
