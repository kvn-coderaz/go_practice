package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kvn-razcode/go_practice/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Println("Listening to port 8080")
	// Start the server
	// http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}