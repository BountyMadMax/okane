package models

import (
	"gorm.io/gorm"
)

// Model of a user.
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
