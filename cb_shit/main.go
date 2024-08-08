package main

import (
	"cb_shit/handler"
	"cb_shit/repository"
	"net/http"
)

func main() {
	repo := repository.NewRepository()
	h := handler.NewHandler(repo)

	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}
