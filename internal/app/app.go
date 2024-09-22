package app

import (
	"context"

	a "github.com/Mubinabd/modestyMart/internal/http"
	"github.com/Mubinabd/modestyMart/internal/http/handlers"
	"github.com/Mubinabd/modestyMart/internal/pkg/config"
	kafka "github.com/Mubinabd/modestyMart/internal/pkg/kafka/consumer"
	prd "github.com/Mubinabd/modestyMart/internal/pkg/kafka/producer"
	s "github.com/Mubinabd/modestyMart/internal/usecase/service"
	"github.com/go-redis/redis/v8"
	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config) {
	// Postgres Connection
	db, err := postgresql.New(cfg)
	if err != nil {
		slog.Error("can't connect to db: %v", err)
		return
	}
	defer db.Db.Close()
	slog.Info("Connected to Postgres")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis_auth:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		slog.Error("Failed to connect to Redis: %v", err)
	}
	slog.Info("Connected to Redis")

	authService := s.NewAuthService(db)
	userService := s.NewUserService(db)

	// Kafka
	brokers := []string{"kafka_auth:9092"}
	cm := kafka.NewKafkaConsumerManager()
	pr, err := prd.NewKafkaProducer(brokers)
	if err != nil {
		slog.Error("Failed to create Kafka producer:", err)
		return
	}

	// HTTP Server
	h := handlers.NewHandler(authService, userService, rdb, &pr)

	router := a.NewGin(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.GRPCPort); err != nil {
		slog.Error("can't start server: %v", err)
	}

	slog.Info("REST server started on port %s", cfg.GRPCPort)
}
