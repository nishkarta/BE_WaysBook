package routes

import (
	"waysbook/handlers"
	"waysbook/pkg/middleware"
	"waysbook/pkg/mysql"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	bookRepository := repositories.RepositoryBook(mysql.DB)
	h := handlers.HandlerBook(bookRepository)

	r.HandleFunc("/books", h.FindBooks).Methods("GET")
	r.HandleFunc("/books-latest", h.LatestBooks).Methods("GET")
	r.HandleFunc("/book/{id}", middleware.Auth(h.GetBookByID)).Methods("GET")
	r.HandleFunc("/book", middleware.Auth(middleware.UploadCover(middleware.UploadPDF(h.AddBook)))).Methods("POST")
	r.HandleFunc("/book/{id}", middleware.Auth(middleware.UploadCover(middleware.UploadPDF(h.UpdateBook)))).Methods("PATCH")
	r.HandleFunc("/book/{id}", middleware.Auth(h.DeleteBook)).Methods("DELETE")

}
