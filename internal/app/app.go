package app

import (
	"context"

	a "github.com/Mubinabd/modestyMart/internal/http"
	"github.com/Mubinabd/modestyMart/internal/http/handlers"
	"github.com/Mubinabd/modestyMart/internal/pkg/config"
	"github.com/Mubinabd/modestyMart/internal/storage/repository"
	prd "github.com/Mubinabd/modestyMart/internal/usecase/kafka"
	s "github.com/Mubinabd/modestyMart/internal/usecase/service"
	"github.com/go-redis/redis/v8"
	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config) {
	// Postgres Connection
	db, err := repository.New(cfg)
	if err != nil {
		slog.Error("can't connect to db: %v", err)
		return
	}
	defer db.Db.Close()
	slog.Info("Connected to Postgres")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		slog.Error("Failed to connect to Redis: %v", err)
	}
	slog.Info("Connected to Redis")

	authService := s.NewAuthService(db)
	userService := s.NewUserService(db)
	productService := s.NewProductService(db)
	paymentService := s.NewPaymentService(db)
	orderService := s.NewOrderService(db)
	categoryService := s.NewCategoryService(db)

	// Kafka
	brokers := []string{"localhost:9092"}
	pr, err := prd.NewKafkaProducer(brokers)
	if err != nil {
		slog.Error("Failed to create Kafka producer:", err)
		return
	}

	// HTTP Server
	h := handlers.NewHandler(productService, paymentService, orderService, categoryService, authService, userService, rdb, &pr)

	router := a.NewGin(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.GRPCPort); err != nil {
		slog.Error("can't start server: %v", err)
	}

	slog.Info("REST server started on port %s", cfg.GRPCPort)
}
