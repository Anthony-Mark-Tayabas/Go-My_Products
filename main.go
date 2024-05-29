package main

import (
	"fmt"
	"my-products/router"
	"my-products/middleware"
	"log"
	"net/http"
)

func main() {
	middleware.InitDB()
	
	r := router.Router()
	http.Handle("/", r)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}