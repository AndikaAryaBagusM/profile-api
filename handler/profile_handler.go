// file: handlers/profile_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/username/profile-api/database"
	"github.com/username/profile-api/models"
)

// GetProfiles: ambil semua profil
func GetProfiles(c *gin.Context) {
	var profiles []models.Profile
	result := database.DB.Find(&profiles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, profiles)
}

// GetProfileByID: ambil satu profil berdasarkan ID
func GetProfileByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var profile models.Profile
	result := database.DB.First(&profile, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// CreateProfile: buat profil baru
func CreateProfile(c *gin.Context) {
	var input struct {
		Name      string `json:"name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Bio       string `json:"bio"`
		AvatarURL string `json:"avatar_url"`
	}
	// Bind JSON payload ke struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := models.Profile{
		Name:      input.Name,
		Email:     input.Email,
		Bio:       input.Bio,
		AvatarURL: input.AvatarURL,
	}
	result := database.DB.Create(&profile)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, profile)
}

// UpdateProfile: perbarui data profil
func UpdateProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var profile models.Profile
	if err := database.DB.First(&profile, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// Struct untuk menerima input update (semi‐partial update)
	var input struct {
		Name      *string `json:"name"`
		Email     *string `json:"email" binding:"omitempty,email"`
		Bio       *string `json:"bio"`
		AvatarURL *string `json:"avatar_url"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update field bila non‐nil
	if input.Name != nil {
		profile.Name = *input.Name
	}
	if input.Email != nil {
		profile.Email = *input.Email
	}
	if input.Bio != nil {
		profile.Bio = *input.Bio
	}
	if input.AvatarURL != nil {
		profile.AvatarURL = *input.AvatarURL
	}

	if err := database.DB.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// DeleteProfile: hapus (soft delete) profil
func DeleteProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var profile models.Profile
	if err := database.DB.First(&profile, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	if err := database.DB.Delete(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent) // 204 No Content
}
