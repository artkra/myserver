package structs

type LineItem struct {
	Name      string `json:"name"`
	ID        int64  `json:"ID"`
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
}
