package service

import (
	"context"

	st "github.com/Mubinabd/modestyMart/internal/storage/repository"
	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
)

type PaymentService struct {
	storage st.Storage
	pb.UnimplementedPaymentsServiceServer
}

func NewPaymentService(storage *st.Storage) *PaymentService {
	return &PaymentService{
		storage: *storage,
	}
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *pb.CreatePaymentReq) (*pb.Void, error) {
	res, err := s.storage.PaymentS.CreatePayment(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PaymentService) ListPayments(ctx context.Context, req *pb.ListPaymentsReq) (*pb.ListPaymentsRes, error) {
	res, err := s.storage.PaymentS.ListPayments(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PaymentService) GetPayment(ctx context.Context, req *pb.GetById) (*pb.Payment, error) {
	res, err := s.storage.PaymentS.GetPayment(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}