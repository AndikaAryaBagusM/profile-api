package models

import (
	"time"

	"gorm.io/gorm"
)

// Profile merepresentasikan tabel profiles di database
type Profile struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"size:100;unique;not null" json:"email"`
	Bio       string         `gorm:"size:255" json:"bio"`        // deskripsi singkat
	AvatarURL string         `gorm:"size:255" json:"avatar_url"` // link gambar profil
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // soft delete
}
