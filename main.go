package main

import (
	"github/maxidelgado/photo-api/handlers"
	"github/maxidelgado/photo-api/utils/router"
	"github/maxidelgado/photo-api/worker"
)

func main() {
	r := router.New()

	// Search returns all the photos with any of the meta fields (author, camera, tags, etc) matching the search term
	r.Engine().Get("/search/:searchTerm", handlers.Search)

	go worker.StartPolling()

	_ = r.Engine().Listen(3000)
}
