package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/service/gateway/handlers"
)

func ServiceRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := handlers.ServiceHandler{}
    r.Get("/", bookHandler.ViewPage)
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}
