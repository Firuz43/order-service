package handlers

import (
	"encoding/json"
	"net/http"

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
func responsdWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

//
