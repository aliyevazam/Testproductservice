package service

import (
	"context"
	"fmt"
	pb "github/FIrstService/template-service/Testproductservice/genproto/product"
	"github/FIrstService/template-service/Testproductservice/genproto/stores"
	l "github/FIrstService/template-service/Testproductservice/pkg/logger"
	"github/FIrstService/template-service/Testproductservice/storage"

	grpcclient "github/FIrstService/template-service/Testproductservice/service/grpc_client"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProductService ...

type ProductService struct {
	store   *grpcclient.ServiceManager
	storage storage.IStorage
	logger  l.Logger
}

// NewProductService ...

func NewProductService(store *grpcclient.ServiceManager, db *sqlx.DB, log l.Logger) *ProductService {
	return &ProductService{
		store:   store,
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.ProductFullInfo) (*pb.ProductFullInfoResponse, error) {
	productReq := &pb.ProductFullInfo{
		Name:       req.Name,
		Model:      req.Model,
		CategoryId: req.CategoryId,
		TypeId:     req.TypeId,
		Price:      req.Price,
		Amount:     req.Amount,
	}
	productResp, err := s.storage.Product().CreateProduct(productReq)
	productInfo := pb.ProductFullInfoResponse{
		Id:         productResp.Id,
		Name:       productResp.Name,
		Model:      productResp.Model,
		CategoryId: productResp.CategoryId,
		TypeId:     productResp.TypeId,
		Price:      productResp.Price,
		Amount:     productResp.Amount,
	}
	fmt.Println(productResp, err)
	if err != nil {
		s.logger.Error("error while creating product full info in product database", l.Any("error creating product full info in product database", err))
		return &pb.ProductFullInfoResponse{}, status.Error(codes.Internal, "something went wrong")
	}
	for _, storeResp := range req.Stores {
		storeReq := stores.StoreRequest{}
		storeReq.Name = storeResp.Name
		for _, addressResp := range storeReq.Address {
			storeReq.Address = append(storeReq.Address, &stores.Address{
				District: addressResp.District,
				Street:   addressResp.Street,
			})
		}
		addressesResp := []*stores.Address{}
		for _, addresStoreInfo := range storeReq.Address {
			addressesResp = append(addressesResp, &stores.Address{
				District: addresStoreInfo.District,
				Street:   addresStoreInfo.Street,
			})
		}
		storeReq.Address = addressesResp
		storeInfo, err := s.store.StoreService().Create(context.Background(), &storeReq)
		if err != nil {
			s.logger.Error("error while creating product full info in store database", l.Any("error creating product ful info in store database", err))
			return &pb.ProductFullInfoResponse{}, status.Error(codes.Internal, "something went wrong")
		}
		addressesRespProduct := []*pb.Address{}
		for _, addresStoreInfo := range storeInfo.Address {
			addressesRespProduct = append(addressesRespProduct, &pb.Address{
				District: addresStoreInfo.District,
				Street:   addresStoreInfo.Street,
			})
		}
		productInfo.Stores = append(productInfo.Stores, &pb.Store{
			Name:      storeInfo.Name,
			Addresses: addressesRespProduct,
		})

	}
	return &productInfo, nil
}

func (s *ProductService) CreateCategory(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	q, err := s.storage.Product().CreateCategory(req)
	if err != nil {
		s.logger.Error("Error while insert", l.Any("Error insert product", err))
		return &pb.Category{}, status.Error(codes.Internal, "something went wrong,please check product infto")
	}
	return q, nil
}

func (s *ProductService) CreateType(ctx context.Context, req *pb.Type) (*pb.Type, error) {
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
