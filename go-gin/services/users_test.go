package services

import (
	"conduit/core"
	"conduit/database"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UsersTestSuite struct {
	suite.Suite
	UserService *UserService
	DB          *gorm.DB
}

func (s *UsersTestSuite) TestValidRegistration() {
	registration := core.UserRegistration{
		Username: "erike",
		Email:    "valid@example.com",
		Password: "abc123",
	}
	registredUser, error := s.UserService.Register(registration)

	s.NotNil(registredUser)
	s.Nil(error)
}

func (s *UsersTestSuite) SetupSuite() {
	log.Println("SetupSuite")
	db, err := gorm.Open("sqlite3", ":memory:")

	db.AutoMigrate(&database.User{})

	s.Require().Nil(err)

	s.UserService = &UserService{DB: db}
	s.DB = db
}

func (s *UsersTestSuite) BeforeTest() {
	s.DB.Begin()
}

func (s *UsersTestSuite) AfterTest() {
	s.DB.Rollback()
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
