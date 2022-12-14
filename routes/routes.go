package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	ProfileRoutes(r)
	AuthRoutes(r)
	ProductRoutes(r)
	ToppingRoutes(r)
}
