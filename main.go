// file: main.go
package main

import (
	"log"

	"github.com/username/profile-api/database"
	"github.com/username/profile-api/routers"
)

func main() {
	// Inisialisasi database & migrasi otomatis
	database.InitDatabase()

	// Setup router
	r := routers.SetupRouter()

	// Jalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
