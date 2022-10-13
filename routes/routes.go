package routes

import "github.com/gorilla/mux"

func RoutesInit(r *mux.Router) {

	AuthRoutes(r)
	LinkRoutes(r)
	SosmedRoutes(r)
}
