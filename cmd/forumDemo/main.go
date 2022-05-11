package main

import (
	"forumDemo/postgres"
	"forumDemo/web"
	"log"
	"net/http"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":3131", h)
}
