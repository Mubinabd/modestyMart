package service

import (
	"context"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	st "github.com/Mubinabd/modestyMart/internal/storage/repository"
)

type CartService struct {
	storage st.Storage
	pb.UnimplementedCartServiceServer
}

func NewCartService(storage *st.Storage) *CartService {
	return &CartService{
		storage: *storage,
	}
}

func (s *CartService) CreateCart(ctx context.Context, req *pb.CreateCartReq) (*pb.Void, error) {
	res, err := s.storage.CartS.CreateCart(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CartService) ListAllCarts(ctx context.Context, req *pb.ListAllCartReq) (*pb.ListAllCartRes, error) {
	res, err := s.storage.CartS.ListAllCarts(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CartService) GetCart(ctx context.Context, req *pb.GetById) (*pb.Cart, error) {
	res, err := s.storage.CartS.GetCart(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
