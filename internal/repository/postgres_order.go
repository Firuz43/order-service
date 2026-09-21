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

// NewPostgresOrderRepository is a constructor returning the interface
func NewPostgresOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &postgresOrderRepository{db: db}
}

func (r *postgresOrderRepository) Update(ctx context.Context, order *models.Order) error {
	//Updating multiple columns
	query := `
		UPDATE orders
		SET customer_name = $1,
			item = $2,
			quantity = $3,
			price = $4,
			status = $5,
			created_at $6,
			updated_at $7,
		WHERE id = $8
	`

	// Exec the update
	result, err := r.db.Exec(ctx, query,
		order.CustomerName, //$1
		order.Item,
		order.Quantity,
		order.Price,
		order.Status,
		order.CreatedAt,
		order.UpdatedAt,
		order.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}

	// Check if any rows were actually updated
	rowsAfected := result.RowsAffected()
	if rowsAfected == 0 {
		return fmt.Errorf("order with ID %s not found", order.ID)
	}

	return nil

	/*
		UPDATE with SET clause changes columns
		WHERE id = $7 = target which row to update
		RowsAffected() = how many rows were changed
		If 0 rows affected, the order didn't exist
	*/
}

func (r *postgresOrderRepository) List(ctx context.Context, limit, offset int) ([]*models.Order, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	query :=

		`SELECT id, customer_name, item, quantity, price, status, created_at, updated_at 
		 FROM orders
		 ORDER BY created_at DESC
		 LIMIT $1 OFFSET $2
		`

	rows, err := r.db.Query(ctx, query, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}
	defer rows.Close()

	orders := make([]*models.Order, 0, limit)

	for rows.Next() {
		order := &models.Order{}

		err := rows.Scan(
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
			return nil, fmt.Errorf("failed to scan order row: %w", err)
		}

		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while literating orders: %w", err)
	}
	return orders, nil
}

// GetByID implements [OrderRepository].
func (r *postgresOrderRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	query := `SELECT id, customer_name, item, quantity, price, status, created_at, updated_at FROM orders WHERE id = $1`

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

// CREATE implements [OrderRepository].
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
