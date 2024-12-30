package models

import "time"

type Note struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Title      string    `gorm:"size:255;not null"`
	Content    string    `gorm:"type:text;not null"`
	CategoryID uint      `gorm:"not null"`              // Relasi ke Category
	Category   Category  `gorm:"foreignKey:CategoryID"` // Relasi ke Category
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
