package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kvn-coderaz/go_practice/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Println("Listening to port 8080")
	// Start the server
	// http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
