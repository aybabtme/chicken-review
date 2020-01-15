package ui

type HTML struct {
	Head Head
	Body Body
}

type Head struct {
	Title       string
	FavIcoURL   string
	Author      string
	Description string
	OG          *OGData
	Twitter     *Twitter
}

type Body struct {
	Content interface{}
}

type OGData struct {
	Title        string
	Type         string
	ImageURL     string
	CanonicalURL string

	SiteName    string
	Description string
}

type Twitter struct {
	ImageURL    string
	Title       string
	Description string
}
