package core

// User ...
type User struct {
	id       string
	email    string
	token    string
	username string
	bio      string
	image    string
}

// UserRegistration ...
type UserRegistration struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Profile ...
type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"string"`
	Following bool   `json:"following"`
}

// Article ...
type Article struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"string"`

	Author Profile
}

// Comment ...
type Comment struct {
}

// Slug ...
type Slug string
