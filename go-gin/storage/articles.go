package storage

import (
	domain "conduit/domain"

	"github.com/jinzhu/gorm"
)

// ArticleStore concrete
type ArticleStore struct {
	db *gorm.DB
}

type articleModel struct {
	gorm.Model
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}

// Create stores a article in the database
func (s *ArticleStore) Create(a *domain.Article) {

}
