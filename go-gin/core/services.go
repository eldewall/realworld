package core

// Articles operations
type Articles interface {
	Feed(u User) ([]Article, error)
	Latest() ([]Article, error)
	Create(a Article) error
	Get(slug string) (Article, error)
	Update(slug string, a Article) (Article, error)
	Delete(slug string) error
}

// Comments ...
type Comments interface {
	Get(a Article) ([]Comment, error)
	Create(a Article, c Comment) (Comment, error)
}

// Favorites
type Favorites interface {
}
