package handlers

import (
	"encoding/json"
	"net/http"
	"ShoppingAppAPI/models"
	"ShoppingAppAPI/utils"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	// Parse request body to get order details
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Process the order (dummy response for demonstration)
	orderResponse := models.OrderResponse{
		Message: "Order placed successfully",
		OrderID: 123,
	}

	// Convert to JSON
	response, err := json.Marshal(orderResponse)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Respond with JSON
	utils.RespondWithJSON(w, http.StatusOK, response)
}

