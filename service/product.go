package service

import (
	"context"
	pb "github/FIrstService/template-service/Testproductservice/genproto/product"
	l "github/FIrstService/template-service/Testproductservice/pkg/logger"
	"github/FIrstService/template-service/Testproductservice/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProductService ...

type ProductService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewProductService ...

func NewProductService(db *sqlx.DB, log l.Logger) *ProductService {
	return &ProductService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	product, err := s.storage.Product().CreateProduct(req)
	if err != nil {
		s.logger.Error("Error while insert", l.Any("Error insert product", err))
		return &pb.Product{}, status.Error(codes.Internal, "something went wrong,please check product infto")
	}
	return product, nil
}

func (s *ProductService) CreateCategory(ctx context.Context,req *pb.Category) (*pb.Category,error) {
	q, err := s.storage.Product().CreateCategory(req)
	if err != nil {
		s.logger.Error("Error while insert", l.Any("Error insert product", err))
		return &pb.Category{}, status.Error(codes.Internal, "something went wrong,please check product infto")
	}
	return q, nil
}

func (s *ProductService) CreateType(ctx context.Context,req *pb.Type) (*pb.Type,error) {
	q, err := s.storage.Product().CreateType(req)
	if err != nil {
		s.logger.Error("Error while insert", l.Any("Error insert product", err))
		return &pb.Type{}, status.Error(codes.Internal, "something went wrong,please check product infto")
	}
	return q, nil
}

func (s *ProductService) GetProductInfoByid(ctx context.Context, req *pb.Ids) (*pb.GetProducts, error) {
	res, err := s.storage.Product().GetProductInfoByid(req)
	if err != nil {
		s.logger.Error("Error while get info", l.Any("Error select product", err))
		return &pb.GetProducts{}, status.Error(codes.Internal, "something went wrong,please check product info")
	}
	return res, err
}

func (s *ProductService) UpdateByid(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	res, err := s.storage.Product().UpdateByid(req)
	if err != nil {
		s.logger.Error("Error while updating", l.Any("Update", err))
		return &pb.Product{}, status.Error(codes.InvalidArgument, "Please recheck user info")
	}
	return res, nil

}

func (s *ProductService) DeleteInfo(ctx context.Context, req *pb.Ids) (*pb.Empty, error) {
	err := s.storage.Product().DeleteInfo(req)
	if err != nil {
		s.logger.Error("Error while delete product", l.Any("Delete", err))
		return &pb.Empty{}, status.Error(codes.Internal, "wrong id for delete")
	}
	return &pb.Empty{}, nil
}
