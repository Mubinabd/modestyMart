package app

import (
	"context"
	"log"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	"github.com/Mubinabd/modestyMart/internal/usecase/service"
	"google.golang.org/protobuf/encoding/protojson"
)

func UserRegister(u *service.AuthService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.RegisterReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.Register(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Register User: %+v",res)
	}
}
func CreateCategory(u *service.CategoryService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateCategoryReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.CreateCategory(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create category: %+v", res)
	}
}
func UpdateCategory(u *service.CategoryService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateCategoryReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.UpdateCategory(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("update category: %+v", res)
	}
}
func CreatePayment(u *service.PaymentService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreatePaymentReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.CreatePayment(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("flash sale product: %+v", res)
	}
}
func CreateOrder(u *service.OrderService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateOrderReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.CreateOrder(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create order: %+v", res)
	}
}
func UpdateOrder(u *service.OrderService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateOrderReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.UpdateOrder(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("update order: %+v", res)
	}
}
func CreateProduct(u *service.ProductService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateProductReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.CreateProduct(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create product: %+v", res)
	}
}
func UpdateProduct(u *service.ProductService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.UpdateProductReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.UpdateProduct(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("update product: %+v", res)
	}
}
func CreateCart(u *service.CartService) func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CreateCartReq
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := u.CreateCart(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("create cart: %+v", res)
	}
}
