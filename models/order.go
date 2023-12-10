package models

type Order struct {
	UserID    int   `json:"userId"`
	ProductID int   `json:"productId"`
	Quantity  int   `json:"quantity"`
}

type OrderResponse struct {
	Message string `json:"message"`
	OrderID int    `json:"orderId"`
}

