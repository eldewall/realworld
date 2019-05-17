package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type ArticleSuite struct {
	suite.Suite
	Storage *ArticleStore
	DB      *gorm.DB
}

func (s *ArticleSuite) SetupSuite() {
	db, err := gorm.Open("sqlite", "test.db")

	if err != nil {
		s.Fail(err.Error())
	}

	s.DB.AutoMigrate(&articleModel{})
}
