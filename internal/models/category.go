package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Notes     []Note    `gorm:"foreignKey:CategoryID"` // Relasi satu ke banyak dengan Note
}
