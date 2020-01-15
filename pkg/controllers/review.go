package controllers

import (
	"chicken-review/pkg/db"
	"chicken-review/pkg/views"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type ReviewController struct {
	db db.DB
}

var showReviewPattern = regexp.MustCompile("^/reviews/([0-9]+)$")

func NewReviewController(db db.DB) *ReviewController {
	return &ReviewController{db: db}
}

func (hdl *ReviewController) Create(resp http.ResponseWriter, req *http.Request) {
	log.Printf("receive request to create a review")
	view := views.NewReviewCreateView(views.DefaultBaseHTMLContext)

	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

func (hdl *ReviewController) Save(resp http.ResponseWriter, req *http.Request) {
	log.Printf("receive request to save a review")
	// TODO: get the posted data
}

func (hdl *ReviewController) Get(resp http.ResponseWriter, req *http.Request) {
	log.Printf("receive request to get a review")
	matches := showReviewPattern.FindStringSubmatch(req.URL.Path)
	if len(matches) != 2 {
		http.Error(resp, "no ID provided", http.StatusBadRequest)
		return
	}
	idStr := matches[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(resp, fmt.Sprintf("ID is not numeric: %v", err.Error()), http.StatusBadRequest)
		return
	}
	review, ok, err := hdl.db.GetReview(id)
	if err != nil {
		log.Printf("finding review: %v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	var view views.View
	if !ok {
		view = views.NewNotFoundView(views.DefaultBaseHTMLContext)
	} else {
		view = views.NewReviewShowView(views.DefaultBaseHTMLContext, review)
	}
	resp.Header().Set("Content-Type", view.ContentType())
	err = view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

func (hdl *ReviewController) List(resp http.ResponseWriter, req *http.Request) {
	log.Printf("receive request to list reviews")
	reviews, err := hdl.db.ListReviews()
	if err != nil {
		log.Printf("listing reviews: %v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	view := views.NewReviewListView(views.DefaultBaseHTMLContext, reviews)
	resp.Header().Set("Content-Type", view.ContentType())
	err = view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

// heejunchicken.com   -> index

// HTTP GET /reviews -> "return a list of reviews" (in HTML)
// HTTP GET /reviews/0 -> "show review number 0" (in HTML)

// HTTP GET /reviews/create -> "show the create review page" (in HTML)
// HTTP POST /reviews/create -> "create a new review, redirect to the the review we just created" (in HTML)
