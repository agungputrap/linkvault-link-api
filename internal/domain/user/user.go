package user

import "time"

type User struct {
	ID        uint     `gorm:"primaryKey"`
	Name      string   `gorm:"not null"`
	Email     Email    `gorm:"unique;not null"`
	Password  Password `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
