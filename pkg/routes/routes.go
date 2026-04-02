package routes

import (
	"github.com/P-kaizoku/BMS/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{book_id}", controllers.DeleteBook).Methods("DELETE")

}
