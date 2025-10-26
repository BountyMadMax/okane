package database

import (
	"fmt"

	"github.com/BountyMadMax/okane/src/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db/okane.db"), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Error opening database: %s", err)
	}

	migrate(db)

	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.UserConfig{},
	)
}
