package routes

import (
	"github.com/gorilla/mux"
	"github.com/kvn-coderaz/go_practice/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	 router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	 router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	 router.HandleFunc("/books/{bookId}", controllers.GetBook).Methods("GET")
	 router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	 router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}