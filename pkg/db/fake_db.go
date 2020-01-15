package db

import (
	"chicken-review/pkg/models"
)

type InMemoryDatabase struct {
	nextID     int
	allReviews []*models.Review
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{}
}

func (db *InMemoryDatabase) CreateReview(newReview *models.Review) (*models.Review, error) {
	id := db.nextID
	db.nextID++
	newReview.ID = id
	db.allReviews = append(db.allReviews, newReview)
	return newReview, nil
}

func (db *InMemoryDatabase) GetReview(id int) (*models.Review, bool, error) {
	for _, review := range db.allReviews {
		if review.ID == id {
			return review, true, nil
		}
	}
	return nil, false, nil
}

func (db *InMemoryDatabase) ListReviews() ([]*models.Review, error) {
	return db.allReviews, nil
}
