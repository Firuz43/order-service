package repository

import (
	"context"
	"fmt"
	"uuid"

	"github.com/Firuz43/order-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresOrderRepository struct {
	db *pgxpool.Pool
}

// Delete implements [OrderRepository].
func (r *postgresOrderRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// List implements [OrderRepository].
func (r *postgresOrderRepository) List(ctx context.Context, limit int, offset int) ([]*models.Order, error) {
	panic("unimplemented")
}

// Update implements [OrderRepository].
func (r *postgresOrderRepository) Update(ctx context.Context, order *models.Order) error {
	panic("unimplemented")
}

// NewPostgresOrderRepository is a constructor returning the interface
func NewPostgresOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &postgresOrderRepository{db: db}
}

func (r *postgresOrderRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	query := `SELECT id, customer_name, item, quantity, price, status, created_at, updated_at FROM ordersWHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)

	var order models.Order
	err := row.Scan(
		&order.ID,
		&order.CustomerName,
		&order.Item,
		&order.Quantity,
		&order.Price,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get order by ID %w", err)
	}

	return &order, nil
}

func (r *postgresOrderRepository) Create(ctx context.Context, order *models.Order) error {
	query := `INSERT INTO orders (id, customer_name, item, quantity, price, status, created_at, updated_at) VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(ctx, query,
		order.ID,
		order.CustomerName,
		order.Item,
		order.Quantity,
		order.Price,
		order.Status,
		order.CreatedAt,
		order.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}
	return nil
}
