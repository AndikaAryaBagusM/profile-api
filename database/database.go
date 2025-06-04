package database

import (
	"log"

	"github.com/AndikaAryaBagusM/profile-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase membuka koneksi ke SQLite dan menjalankan auto‐migrate
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("profiles.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// Auto‐migrate: buat tabel berdasarkan struct di models.Profile
	err = DB.AutoMigrate(&models.Profile{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
}
