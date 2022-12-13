package main

import (
	"fmt"
	"net/http"
	"os"
	"waysbook/database"
	"waysbook/pkg/mysql"
	"waysbook/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
)

func main() {
// 	errEnv := godotenv.Load()
// 	if errEnv != nil {
// 		panic("Failed to load env file")
// 	}

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")

	fmt.Println("server running localhost:5000")
	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
