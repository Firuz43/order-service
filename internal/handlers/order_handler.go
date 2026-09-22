package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Firuz43/order-service/internal/models"
	"github.com/Firuz43/order-service/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

// GET /orders/{id}
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	//Extract URL Parameter
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	order, err := h.svc.GetOrder(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Order not found")
	}

	respondWithJSON(w, http.StatusOK, order)
}

// GET /orders?limit10&offset=0
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	orders, err := h.svc.ListOrders(r.Context(), limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to list orders")
		return
	}
	respondWithJSON(w, http.StatusOK, orders)
}
