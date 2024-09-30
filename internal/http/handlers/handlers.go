package handlers

import (
	kafka "github.com/Mubinabd/modestyMart/internal/usecase/kafka"
	s "github.com/Mubinabd/modestyMart/internal/usecase/service"
	"github.com/go-redis/redis/v8"
)

type Handlers struct {
	Product  *s.ProductService
	Payment  *s.PaymentService
	Order    *s.OrderService
	Category *s.CategoryService
	Auth     *s.AuthService
	User     *s.UserService
	Cart     *s.CartService
	RDB      *redis.Client
	Producer kafka.KafkaProducer
}

func NewHandler(product *s.ProductService, payment *s.PaymentService, order *s.OrderService, category *s.CategoryService, auth *s.AuthService, user *s.UserService, cart *s.CartService, rdb *redis.Client, pr *kafka.KafkaProducer) *Handlers {
	return &Handlers{
		Product:  product,
		Payment:  payment,
		Order:    order,
		Category: category,
		Auth:     auth,
		User:     user,
		Cart:     cart,
		RDB:      rdb,
		Producer: *pr,
	}
}
