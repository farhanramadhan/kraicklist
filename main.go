package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"challenge.haraj.com.sa/kraicklist/cmd/serve"
	"challenge.haraj.com.sa/kraicklist/internal/searchs/repository"
	"challenge.haraj.com.sa/kraicklist/internal/services"
)

func main() {
	// initialize searcher
	repoSearch := repository.NewSearchRepository()

	svcSearch := services.NewSearchService(repoSearch)

	handlerSearch := serve.NewRouteHandler(svcSearch)

	err := svcSearch.Load(context.TODO(), "data.gz")
	if err != nil {
		log.Fatalf("unable to load search data due: %v", err)
	}

	// define http handlers
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/search", handlerSearch)
	// define port, we need to set it as env for Heroku deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	// start server
	fmt.Printf("Server is listening on %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}
