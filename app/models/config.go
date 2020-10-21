package models

// TODO: 1. Finish models
type LineItem struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Creatives []*Creative
}

type Creative struct {
	ID   int64
	Name string
	Type string
	URL  string
}

type Config struct {
	LineItems []*LineItem
	Creatives []*Creative
}
