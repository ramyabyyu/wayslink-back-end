package routes

import (
	"wayslink/handlers"
	"wayslink/pkg/middlewares"
	"wayslink/pkg/postgre"
	"wayslink/repositories"

	"github.com/gorilla/mux"
)

func SosmedRoutes(r *mux.Router) {
	sosmedRepository := repositories.RepositorySosmed(postgre.DB)
	h := handlers.HandlerSosmed(sosmedRepository)

	r.HandleFunc("/sosmed", middlewares.Auth(middlewares.UploadFile(h.CreateSosmed))).Methods("POST")
	r.HandleFunc("/sosmed/{link_id}", middlewares.Auth(h.FindSosmedsByLinkID)).Methods("GET")
}
