package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the SQLite connection
func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("crud.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(" Failed to connect to database:", err)
	}

	fmt.Println(" Connected to SQLite database successfully")
}