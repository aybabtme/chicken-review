package models

type Review struct {
	ID                int
	Title             string
	Author            string
	DefaultPictureURL string
	PictureURLs       []string
	StoreName         string
	Date              string
	PhoneNumber       string
	Comment           string
	Score             int
}
