package order

import (
	"errors"
	"time"
)

var (
	ErrCustomerRequired = errors.New("customer is mandatory")
	ErrNotFound         = errors.New("order not found")
)

type Status string

const StatusPending Status = "pending"

// entity
type Order struct {
	ID         string    `json:"id"`
	Status     Status    `json:"status"`
	CustomerID string    `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}

// request
type OrderItemInput struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type PlaceOrderRequest struct {
	CustomerID string           `json:"customer_id"`
	Items      []OrderItemInput `json:"items"`
}
