package controllers

import (
	"chicken-review/pkg/db"
	"chicken-review/pkg/views"
	"log"
	"net/http"
)

type ReviewController struct {
	db db.DB
}

func NewReviewController(db db.DB) *ReviewController {
	return &ReviewController{db: db}
}

func (hdl *ReviewController) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	var view views.View

	switch req.URL.Path {
	case "/reviews", "/reviews/":
		log.Printf("receive request to list reviews")

		// get a list of reviews
		reviews, err := hdl.db.ListReviews()
		if err != nil {
			log.Printf("listing reviews: %v", err)
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}
		// render it in HTML
		view = &views.ReviewListView{Reviews: reviews}

	default:
		view = views.NotFoundView{}
	}

	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

// heejunchicken.com   -> index

// HTTP GET /reviews -> "return a list of reviews" (in HTML)
// HTTP GET /reviews/0 -> "show review number 0" (in HTML)

// HTTP GET /reviews/create -> "show the create review page" (in HTML)
// HTTP POST /reviews/create -> "create a new review, redirect to the the review we just created" (in HTML)
