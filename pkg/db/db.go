package db

import "chicken-review/pkg/models"

type DB interface {
	CreateReview(newReview *models.Review) (*models.Review, error)
	GetReview(id int) (*models.Review, bool, error)
	ListReviews() ([]*models.Review, error)
}
