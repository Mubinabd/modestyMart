package service

import (
	"context"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	st "github.com/Mubinabd/modestyMart/internal/storage/repository"
)

type CategoryService struct {
	storage st.Storage
	pb.UnimplementedCategoriesServiceServer
}

func NewCategoryService(storage *st.Storage) *CategoryService {
	return &CategoryService{
		storage: *storage,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryReq) (*pb.Void, error) {
	res, err := s.storage.CategoryS.CreateCategory(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryService) ListCategories(ctx context.Context, req *pb.ListAllCategoriesReq) (*pb.ListCategoriesRes, error) {
	res, err := s.storage.CategoryS.ListCategories(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryReq) (*pb.Void, error) {
	res, err := s.storage.CategoryS.UpdateCategory(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.CategoryS.DeleteCategory(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, req *pb.GetById) (*pb.Categories, error) {
	res, err := s.storage.CategoryS.GetCategory(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
