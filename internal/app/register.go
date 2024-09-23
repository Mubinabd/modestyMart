
package app

import (
	"errors"

	"github.com/Mubinabd/modestyMart/internal/pkg/config"
	"github.com/Mubinabd/modestyMart/internal/usecase/kafka"
)

func Register(h *KafkaHandler, cfg *config.Config) error {

	brokers := []string{cfg.KafkaUrl}
	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "create", "create-id", h.Register()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update", "update-id", h.EditProfile()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "edit", "edit", h.EditSetting()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'edit' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	
	if err := kcm.RegisterConsumer(brokers, "create-cat", "create-cat-id", h.CreateCategory()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-cat' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update-cat", "update-cat-id", h.UpdateCategory()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-cat' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-pay", "create-pay-id", h.CreatePayment()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-pay' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-order", "order-id", h.CreateOrder()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'order' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update-order", "update-order-id", h.UpdateOrder()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-order' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "create-product", "create-product-id", h.CreateProduct()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'create-product' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "update-product", "update-product-id", h.UpdateProduct()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'update-product' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	return nil
}
