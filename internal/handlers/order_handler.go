package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Firuz43/order-service/internal/models"
	"github.com/Firuz43/order-service/internal/service"
)

type OrderHandler struct {
	svc service.OrderService //injected service dependecy here
}

func NewOrderHandler(svc service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

// JSON HELPER FUNCTION
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to encode JSON response"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// JSON Error Helper Function
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// POST /orders
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req models.CreateOrderRequest

	// Decode JSON Body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Call Service
	order, err := h.svc.CreateOrder(r.Context(), &req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidInput) {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to create order")
		return
	}

	respondWithJSON(w, http.StatusCreated, order)
}
