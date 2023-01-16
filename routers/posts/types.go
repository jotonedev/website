package posts

import (
	"encoding/xml"
	"jotone.eu/database"
)

type URLSet struct {
	XMLName xml.Name        `xml:"urlset"`
	XMLNS   string          `xml:"xmlns,attr"`
	Url     []database.Post `xml:"url"`
}
