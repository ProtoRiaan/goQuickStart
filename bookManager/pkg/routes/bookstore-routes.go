package routes

import (
	"github.com/ProtoRiaan/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookBookByID).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookBookByID).Methods("DELETE")
}
