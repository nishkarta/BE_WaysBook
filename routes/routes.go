package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	BookRoutes(r)
	CartRoutes(r)

	AuthRoutes(r)
}
