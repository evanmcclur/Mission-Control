package models

import (
	"time"
)

// Author model
type Author struct {
	// Author first name
	FirstName string
	// Author last name
	LastName string
}

// Article model
type Article struct {
	// The title of the article
	Title string
	// The article author
	Author *Author
	// The article strapline (if applicable)
	StrapLine string
	// When the article was published
	Published time.Time
	// If this article is archived it will persist in the db
	Archived bool
}