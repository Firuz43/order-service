package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusCompleted OrderStatus = "COMPLETED"
	StatusCancelled OrderStatus = "CANCELLED"
)

type Order struct {
	ID           uuid.UUID   `json:"id"`
	CustomerName string      `json:"customer_name"`
	Item         string      `json:"item"`
	Quantity     int         `json:"quantity"`
	Price        float64     `json:"price"`
	Status       OrderStatus `json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// Payload expected from HTTP client
type CreateOrderRequest struct {
	CustomerName string  `json:"customer_name"`
	Item         string  `json:"item"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json: "price"`
}
