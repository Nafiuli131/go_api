package migrate

import (
	"log"

	"github.com/nafiul/api_tutorial/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	// Auto-migrate all models
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal(" Failed to migrate database: ", err)
	}
	log.Println(" Database migrated successfully")
}
