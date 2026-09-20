package repository

// import (
// 	"context"

// 	"github.com/Firuz43/order-service/internal/models"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type postgresOrderRepository struct {
// 	db *pgxpool.Pool
// }

// // NewPostgresOrderRepository is a constructor returning the interface
// func NewPostgresOrderRepository(db *pgxpool.Pool) OrderRepository {
// 	return &postgresOrderRepository{db: db}
// }

// func (r *postgresOrderRepository) Create(ctx context.Context, order *models.Order) error {
// 	query := `INSERT INTO orders (id, customer_name, item, quantity, price, status, created_at, updated_at) VALUES
// 	($1, $2, $3, $4, $5, $6, $7, $8)`
// 	_, err := r.db.Exec(ctx, query,

// 	)
// }
