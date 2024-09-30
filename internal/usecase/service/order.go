package service

import (
	"context"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	st "github.com/Mubinabd/modestyMart/internal/storage/repository"
)

type OrderService struct {
	storage st.Storage
	pb.UnimplementedOrdersServiceServer
}

func NewOrderService(storage *st.Storage) *OrderService {
	return &OrderService{
		storage: *storage,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.Void, error) {
	res, err := s.storage.OrderS.CreateOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) ListAllOrders(ctx context.Context, req *pb.ListOrdersReq) (*pb.ListOrdersRes, error) {
	res, err := s.storage.OrderS.ListAllOrders(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderReq) (*pb.Void, error) {
	res, err := s.storage.OrderS.UpdateOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.OrderS.DeleteOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetById) (*pb.Orders, error) {
	res, err := s.storage.OrderS.GetOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) GetOrderByProductID(ctx context.Context, req *pb.OrderByProductId) (*pb.GetOrdersRes, error) {
	res, err := s.storage.OrderS.GetOrderByProductID(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
