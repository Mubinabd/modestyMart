package app

import (
	"errors"
	"log"

	"github.com/Mubinabd/modestyMart/internal/usecase/kafka"
	"github.com/Mubinabd/modestyMart/internal/usecase/service"
)

func Register(brokers []string, kcm *kafka.KafkaConsumerManager, authService *service.AuthService, cartService *service.CartService, orderService *service.OrderService, categoryService *service.CategoryService, paymentService *service.PaymentService, productService *service.ProductService, notificationService *service.NotificationService) error {

	if err := kcm.RegisterConsumer(brokers, "create", "create-id", UserRegister(authService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}

	if err := kcm.RegisterConsumer(brokers, "create-cat", "create-cat-id", CreateCategory(categoryService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-cat' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update-cat", "update-cat-id", UpdateCategory(categoryService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-cat' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-pay", "create-pay-id", CreatePayment(paymentService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-pay' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-order", "order-id", CreateOrder(orderService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'order' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update-order", "update-order-id", UpdateOrder(orderService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-order' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-product", "create-product-id", CreateProduct(productService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-product' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update-product", "update-product-id", UpdateProduct(productService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-product' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-cart", "create-cart-id", CreateCart(cartService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-cart' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}

	if err := kcm.RegisterConsumer(brokers, "notification-create", "notification", NotificationCreateHandler(notificationService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'notification-create' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	if err := kcm.RegisterConsumer(brokers, "notify-all", "notification-all", NotifyAllHandler(notificationService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'notify-all' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}

	return nil
}
