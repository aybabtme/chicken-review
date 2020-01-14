package controllers

import (
	"chicken-review/pkg/views"
	"log"
	"net/http"
)

type NotFoundController struct {
}

// HTTP GET index -> REDIRECT /reviews

func (hdl *NotFoundController) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	view := views.NotFoundView{}
	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}
