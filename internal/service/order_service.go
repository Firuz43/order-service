package service

import (
	"context"
	"errors"

	"github.com/Firuz43/order-service/internal/models"
	"github.com/Firuz43/order-service/internal/repository"
	"github.com/google/uuid"
)

// Common business errors that can be returned by the service layer
var (
	ErrInvalidInput = errors.New("invalid input data")
	ErrorNotFound   = errors.New("order not found")
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *models.CreateOrderRequest) (*models.Order, error)
	GetOrder(ctx context.Context, id uuid.UUID) (*models.Order, error)
	ListOrders(ctx context.Context, limit, offset int) ([]*models.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

// NewOrderService is a constructor for dependency injection
func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}
