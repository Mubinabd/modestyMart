package service

import (
	"context"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	st "github.com/Mubinabd/modestyMart/internal/storage/repository"
	"github.com/Mubinabd/modestyMart/internal/usecase/kafka"
	"github.com/Mubinabd/modestyMart/internal/usecase/minio"
)

type ProductService struct {
	storage  st.Storage
	producer kafka.KafkaProducer
	minio    *minio.MinIO
	pb.UnimplementedProductsServiceServer
}

func NewProductService(storage *st.Storage, minio *minio.MinIO, producer kafka.KafkaProducer) *ProductService {
	return &ProductService{
		minio:    minio,
		producer: producer,
		storage:  *storage,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (*pb.Void, error) {
	res, err := s.storage.ProductS.CreateProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductService) ListAllProducts(ctx context.Context, req *pb.ListAllProductsReq) (*pb.ListAllProductsRes, error) {
	res, err := s.storage.ProductS.ListAllProducts(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductReq) (*pb.Void, error) {
	res, err := s.storage.ProductS.UpdateProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.ProductS.DeleteProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetById) (*pb.Products, error) {
	res, err := s.storage.ProductS.GetProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *ProductService) GetCategory(ctx context.Context, req *pb.GetCategoryReq) (*pb.GetCategoryRes, error) {
	res, err := s.storage.ProductS.GetCategory(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductService) GetProductsByPriceRange(ctx context.Context, req *pb.GetProductsByPriceRangeReq) (*pb.ListAllProductsRes, error) {
	res, err := s.storage.ProductS.GetProductsByPriceRange(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
