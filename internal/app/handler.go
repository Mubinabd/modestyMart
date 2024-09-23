package app

import (
	"context"
	"log"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	"github.com/Mubinabd/modestyMart/internal/usecase/service"
	"google.golang.org/protobuf/encoding/protojson"
)

type KafkaHandler struct {
	auth *service.AuthService
	user *service.UserService
	product *service.ProductService
	order *service.OrderService
	payment *service.PaymentService
	category *service.CategoryService
}

func (h *KafkaHandler) Register() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.RegisterReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.auth.Register(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Register User: %+v", res)
	}
}
func (h *KafkaHandler) EditProfile() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UserRes
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.user.EditProfile(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Edit profile: %+v", res)
	}
}

func (h *KafkaHandler) EditSetting() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.SettingReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.user.EditSetting(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Edit Setting: %+v", res)
	}
}

func (h *KafkaHandler) CreateCategory() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateCategoryReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.category.CreateCategory(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create category: %+v", res)
	}
}
func (h *KafkaHandler) UpdateCategory() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateCategoryReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.category.UpdateCategory(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("update category: %+v", res)
	}
}
func (h *KafkaHandler) CreatePayment() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreatePaymentReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.payment.CreatePayment(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("flash sale product: %+v", res)
	}
}
func (h *KafkaHandler) CreateOrder() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateOrderReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.order.CreateOrder(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create order: %+v", res)
	}
}
func (h *KafkaHandler) UpdateOrder() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateOrderReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.order.UpdateOrder(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("update order: %+v", res)
	}
}
func (h *KafkaHandler) CreateProduct() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateProductReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.product.CreateProduct(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create product: %+v", res)
	}
}
func (h *KafkaHandler) UpdateProduct() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateProductReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.product.UpdateProduct(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("update product: %+v", res)
	}
}