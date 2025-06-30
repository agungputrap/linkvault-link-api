package link

import (
	"github.com/lib/pq"
	"time"
)

type Link struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Url         string `gorm:"not null"`
	Description string
	Tags        pq.StringArray `gorm:"type:text[]"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
