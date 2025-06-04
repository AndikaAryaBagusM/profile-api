// file: routers/router.go
package routers

import (
	"github.com/AndikaAryaBagusM/profile-api/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRouter: inisialisasi Gin dan daftarkan route‐route terkait profile
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Grouping endpoint terkait profiles
	profiles := r.Group("/profiles")
	{
		profiles.GET("", handlers.GetProfiles)          // GET /profiles
		profiles.GET("/:id", handlers.GetProfileByID)   // GET /profiles/:id
		profiles.POST("", handlers.CreateProfile)       // POST /profiles
		profiles.PUT("/:id", handlers.UpdateProfile)    // PUT /profiles/:id
		profiles.DELETE("/:id", handlers.DeleteProfile) // DELETE /profiles/:id
	}

	return r
}
