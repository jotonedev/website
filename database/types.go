package database

import (
	"encoding/xml"
	"time"
)

type Post struct {
	XMLName xml.Name `xml:"url"`

	ID          int       `xml:"-"`
	Title       string    `xml:"-"`
	Description string    `xml:"-"`
	CreatedAt   time.Time `xml:"-"`
	Content     string    `xml:"-"`
	Lang        string    `xml:"-"`
	Thumbnail   Thumbnail `xml:"-"`
	Keywords    string    `xml:"-"`

	// Used for sitemap
	URL       string    `xml:"loc"` // This is not in the database, but is used for the sitemap
	UpdatedAt time.Time `xml:"lastmod"`
}

type Thumbnail struct {
	ID     int    `xml:"id,omitempty"`
	Alt    string `xml:"alt_text,omitempty"`
	URL    string `xml:"image,omitempty"`
	Height int    `xml:"height,omitempty"`
	Width  int    `xml:"width,omitempty"`
	Type   string `xml:"type,omitempty"`
}
