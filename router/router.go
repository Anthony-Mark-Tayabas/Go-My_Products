package router

import (
	"my-products/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	
	
	router.HandleFunc("/products", middleware.CreateProduct).Methods("POST")
	router.HandleFunc("/products", middleware.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{name}", middleware.GetProduct).Methods("GET")
	router.HandleFunc("/products/{name}", middleware.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{name}", middleware.DeleteProduct).Methods("DELETE")
	
	return router
}