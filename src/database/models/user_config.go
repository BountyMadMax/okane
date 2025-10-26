package models

// Holds config of an user.
type UserConfig struct {
	Key    string `gorm:"primaryKey"`
	Value  string
	UserID uint
	User   User
}
