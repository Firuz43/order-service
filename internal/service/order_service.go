package service

import (
	"context"
	"errors"
	"fmt"
	"time"

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

func (s *orderService) CreateOrder(ctx context.Context, req *models.CreateOrderRequest) (*models.Order, error) {
	// 1. Business Logic Validations
	if req.CustomerName == "" || req.Item == "" {
		return nil, fmt.Errorf("%w: customer name and item are required", ErrInvalidInput)
	}
	if req.Quantity <= 0 || req.Price <= 0 {
		return nil, fmt.Errorf("%w: quantity and price must be greater than zero", ErrInvalidInput)
	}

	// 2. Domain Entity Construction
	now := time.Now().UTC()
	order := &models.Order{
		ID:           uuid.New(),
		CustomerName: req.CustomerName,
		Item:         req.Item,
		Quantity:     req.Quantity,
		Price:        req.Price,
		Status:       models.StatusPending,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// 3. Persist via Repository
	if err := s.repo.Create(ctx, order); err != nil {
		return nil, fmt.Errorf("service failed to create order: %w", err)
	}

	// LATER WE WILL PUBLISH A RABBITMQ EVENT HERE

	return order, nil
}

func (s *orderService) GetOrder(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *orderService) ListOrders(ctx context.Context, limit, offset int) ([]*models.Order, error) {
	return s.repo.List(ctx, limit, offset)
}
