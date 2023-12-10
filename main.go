package main

import (
	"log"
	"net/http"
	"ShoppingAppAPI/handlers"
)

func main() {
	// Define routes
	http.HandleFunc("/products", handlers.GetProducts)
	http.HandleFunc("/orders", handlers.PlaceOrder)

	// Start the server
	port := ":8080"
	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
