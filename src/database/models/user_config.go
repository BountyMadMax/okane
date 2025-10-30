package models

import (
	"gorm.io/gorm"
)

// Holds the config of a User.
type UserConfig struct {
	gorm.Model
	UserID uint
	User   User
	Theme  string `json:"theme"`
}
