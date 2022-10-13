package routes

import (
	"wayslink/handlers"
	"wayslink/pkg/middlewares"
	"wayslink/pkg/postgre"
	"wayslink/repositories"

	"github.com/gorilla/mux"
)

func LinkRoutes(r *mux.Router) {
	linkRepository := repositories.RepositoryLink(postgre.DB)
	h := handlers.HandlerLink(linkRepository)

	r.HandleFunc("/link", middlewares.Auth(middlewares.UploadFile(h.CreateLink))).Methods("POST")
	r.HandleFunc("/links", middlewares.Auth(h.FindLinks)).Methods("GET")
	r.HandleFunc("/preview_link/{unique_link}", middlewares.Auth(h.PreviewLink)).Methods("GET")
}
