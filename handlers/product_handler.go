package handlers

import (
	"encoding/json"
	"net/http"
	"ShoppingAppAPI/models"
	"ShoppingAppAPI/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Fetch products (dummy data for demonstration)
	products := []models.Product{
		{ID: 1, Name: "Product 1", Price: 29.99},
		{ID: 2, Name: "Product 2", Price: 39.99},
	}

	// Convert to JSON
	response, err := json.Marshal(products)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Respond with JSON
	utils.RespondWithJSON(w, http.StatusOK, response)
}
