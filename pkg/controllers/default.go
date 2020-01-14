package controllers

import (
	"log"
	"net/http"
)

type DefaultController struct {
}

// HTTP GET index -> REDIRECT /reviews

func (hdl *DefaultController) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/reviews", http.StatusPermanentRedirect)
	log.Printf("defaultController: serving redirect")
}
