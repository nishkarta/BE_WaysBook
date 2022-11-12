package routes

import (
	"waysbook/handlers"
	"waysbook/pkg/middleware"
	"waysbook/pkg/mysql"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart/add/{bookID}", middleware.Auth(h.AddToCart)).Methods("POST")
	r.HandleFunc("/carts/{userID}", middleware.Auth(h.GetCartsByUser)).Methods("GET")
	r.HandleFunc("/carts", middleware.Auth(h.FindCarts)).Methods("GET")
	r.HandleFunc("/current-carts", middleware.Auth(h.GetCartsByCurrentUser)).Methods("GET")
	r.HandleFunc("/cart/{cartID}", middleware.Auth(h.GetCartByID)).Methods("GET")
	r.HandleFunc("/cart/delete/{cartID}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
}
