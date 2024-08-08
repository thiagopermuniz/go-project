package handler

import (
	"cb_shit/repository"
	"log"
	"net/http"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.GetData()
	if err != nil {
		log.Println("Error fetching data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if data == "no content" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write([]byte(data))
}
