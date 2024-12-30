package config

import (
	"github.com/romanmufid16/notes_api_backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func DatabaseConnection() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Jika query lambat lebih dari 1 detik
			LogLevel:      logger.Info, // Menampilkan log pada level Info
			Colorful:      true,        // Menampilkan log berwarna
		},
	)

	dsn := os.Getenv("DATABASE_URI")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func SyncDatabase() {
	DB.AutoMigrate(&models.Category{}, &models.Note{})
}
