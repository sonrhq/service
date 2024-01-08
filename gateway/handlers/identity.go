package handlers

import (
	"net/http"

	"github.com/sonrhq/service/gateway/templates/components"
)

type ServiceHandler struct {
}

func (b ServiceHandler) ViewPage(w http.ResponseWriter, r *http.Request) {
	err := components.Page(0,1).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (b ServiceHandler) ListBooks(w http.ResponseWriter, r *http.Request)  {}
func (b ServiceHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b ServiceHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b ServiceHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b ServiceHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
