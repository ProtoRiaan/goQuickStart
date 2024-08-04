package routes

import (
	"github.com/gorilla/mux"
	"github.com/protoriaan/go-bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookBookByID).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookBookByID).Methods("DELETE")
}
