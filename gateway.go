package service

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/service/gateway/routes"
)

func RegisterGateway(mux *chi.Mux) {
    mux.Mount("/service", routes.ServiceRoutes())
}
