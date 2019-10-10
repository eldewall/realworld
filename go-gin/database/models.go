package database

import "github.com/jinzhu/gorm"

// User in the database domain
type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	Email    string `gorm:"unique_index"`
}
