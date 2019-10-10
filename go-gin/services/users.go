package services

import (
	"conduit/core"

	"github.com/jinzhu/gorm"
)

// UserService ...
type UserService struct {
	DB *gorm.DB
}

// Register a new User
func (us *UserService) Register(registration core.UserRegistration) (core.User, error) {
	return core.User{}, nil
}
