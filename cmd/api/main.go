package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Firuz43/order-service/internal/config"
	"github.com/Firuz43/order-service/internal/db"
	"github.com/Firuz43/order-service/internal/handlers"
	"github.com/Firuz43/order-service/internal/repository"
	"github.com/Firuz43/order-service/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Conenct to PostgreSQL
	dbpool, err := db.ConnectPostgres(cfg)

	if err != nil {
		log.Fatalf("Fatal: Could not initialize database %v", err)
	}

	// Clean up pool when main function exists
	defer dbpool.Close()

	log.Printf("Server starting on port %s", cfg.ServerPort)

	// Next: We will set up HTTP server router & route here

	// 3. DEPENDENCY INJECTION HERE
	repo := repository.NewPostgresOrderRepository(dbpool)
	svc := service.NewOrderService(repo)
	handler := handlers.NewOrderHandler(svc)

	// 4. Router Setup WIth Essentioal Middlewares
	r := chi.NewRouter()

	// Built-in Chi Middlewares
	r.Use(middleware.RequestID)                 // Attaches unique Request ID to context
	r.Use(middleware.RealIP)                    // Obtains user IP behind load balancers
	r.Use(middleware.Logger)                    // Logs request duration, path, and HTTP status
	r.Use(middleware.Recoverer)                 // Catches panics and prevents server crashes
	r.Use(middleware.Timeout(60 * time.Second)) // Auto-cancels context after 60s

	// 5. Route Declarations
	r.Route("/orders", func(r chi.Router) {
		r.Post("/", handler.CreateOrder)
		r.Get("/", handler.ListOrders)
		r.Get("/{id}", handler.GetOrder)
	})

	// Health check endpoint for Docker & KUBER
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"UP"}`))
	})

	//6. Start SErver
	log.Printf("Server listening on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("Server stopped unexpectedly: %v", err)
	}
}
