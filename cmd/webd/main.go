package main

import (
	"chicken-review/pkg/controllers"
	"chicken-review/pkg/db"
	"chicken-review/pkg/models"
	"chicken-review/pkg/router"
	"log"
	"net"
	"net/http"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	myDB := db.NewInMemoryDatabase()

	review1 := &models.Review{
		StoreName:         "Nonhyeon Chicken",
		Author: "Antoine",
		Date:              "2019-12-20",
		Title:             "Delicious And Crispy",
		DefaultPictureURL: "https://assets-aybabt-me.sfo2.cdn.digitaloceanspaces.com/boat/lifepo_project/v0/lifepo4_upgrade_v0.png",
		PictureURLs: []string{
			"https://assets-aybabt-me.sfo2.cdn.digitaloceanspaces.com/boat/lifepo_project/v0/lifepo4_upgrade_v0.png",
		},
	}
	review2 := &models.Review{
		StoreName: "BBQ Chicken",
		Author: "Antoine",
		Date:      "2020-01-13",
		Title:     "What A Disappointment!",
	}

	myDB.CreateReview(review1)
	myDB.CreateReview(review2)

	defaultCtrl := &controllers.DefaultController{}
	notFoundCtrl := &controllers.NotFoundController{}
	reviewsCtrl := controllers.NewReviewController(myDB)

	rtr := router.NewRouter(notFoundCtrl)
	rtr.AddRule("reviews", "GET", "^/reviews/create$", reviewsCtrl.Create)
	rtr.AddRule("reviews", "GET", "^/reviews/([0-9]+)$", reviewsCtrl.Get)
	rtr.AddRule("reviews", "GET", "^/reviews/?$", reviewsCtrl.List)
	rtr.AddRule("default", "GET", "^/$", defaultCtrl.ServeHTTP)

	srv := http.Server{Handler: rtr}

	log.Printf("starting server at address %q", l.Addr().String())
	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
