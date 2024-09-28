package service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	st "github.com/Mubinabd/modestyMart/internal/storage/repository"
	"github.com/Mubinabd/modestyMart/internal/usecase/kafka"
	"github.com/Mubinabd/modestyMart/internal/usecase/minio"
	pdfmaker "github.com/Mubinabd/modestyMart/internal/usecase/pdf_maker"
	"google.golang.org/protobuf/encoding/protojson"
)

type ProductService struct {
	storage  st.Storage
	producer kafka.KafkaProducer
	minio    *minio.MinIO
	pb.UnimplementedProductsServiceServer
}

func NewProductService(storage *st.Storage, minio *minio.MinIO,producer kafka.KafkaProducer) *ProductService {
	return &ProductService{
		minio:   minio,
		producer: producer,
		storage: *storage,
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
	products, err := s.storage.ProductS.GetProduct(&pb.GetById{Id: req.Id})
	if err != nil {
		return nil, errors.New("error: " + err.Error())
	}
	if req.Body.ImageUrl == "pending" {
		res, err := pdfmaker.GenerateCertificate(
			products.Name,
			products.ImageUrl,
			"ModestyMart",
			s.minio.Cf,
		)
		if err != nil {
			return nil, errors.New("internal - usecases- service -product -genrating certificate error: " + err.Error())
		}
		cer_url, err := s.minio.Upload(*res, ".pdf")
		if err != nil {
			return nil, errors.New("internal usecases minio upload error: " + err.Error())
		}
		req.Body.ImageUrl = *cer_url
	}
	// write to kafka
	message := fmt.Sprintf("dear %s your certificate request has been %s", products.Name, req.Body.ImageUrl)
	notification := pb.NotificationCreate{
		RecieverId: products.Id,
		Message:    message,
	}
	input, err := protojson.Marshal(&notification)
	if err != nil {
		return nil, errors.New("internal usecases protojson marshal error: " + err.Error())
	}
	err = s.producer.ProduceMessages("notification-create", input)
	if err != nil {
		return nil, errors.New("internal usecases kafka produce error: " + err.Error())
	}

	return s.storage.ProductS.UpdateProduct(req)

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
