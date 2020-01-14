package models

type Review struct {
	ID          int
	Title       string
	PictureURLs []string
	StoreName   string
	Date        string
	PhoneNumber string
	Comment     string
	Score       int
}
