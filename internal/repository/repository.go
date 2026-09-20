package repository

import (
	"context"
	"uuid"

	"github.com/Firuz43/order-service/internal/models"
)

type OrderRepository interface {
	// CreateOrder creates a new order in the database
	Create(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Order, error)
	List(ctx context.Context, limit, offset int)([]*models.Order, error)
	Update(ctx context.Context, order *models.Order)error
	Delete(ctx context.Context, id uuid.UUID) error
}
