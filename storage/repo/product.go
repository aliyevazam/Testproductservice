package repo

import (
	pb "github/FIrstService/template-service/Testproductservice/genproto/product"
)

// UserStorageI ...
type ProductStorageI interface {
	CreateProduct(*pb.Product) (*pb.Product, error)
	CreateCategory(*pb.Category) (*pb.Category, error)
	CreateType(*pb.Type) (*pb.Type, error)
	GetProductInfoByid(*pb.Ids) (*pb.GetProducts, error)
	DeleteInfo(*pb.Ids) error
	UpdateByid(*pb.Product) (*pb.Product, error)
}
