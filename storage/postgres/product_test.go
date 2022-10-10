package postgres

import (
	"github/FIrstService/template-service/Testproductservice/config"
	pb "github/FIrstService/template-service/Testproductservice/genproto/product"
	"github/FIrstService/template-service/Testproductservice/pkg/db"
	"github/FIrstService/template-service/Testproductservice/storage/repo"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProductSuiteTest struct {
	suite.Suite
	ClenUpFunc func()
	Repository repo.ProductStorageI
}

func (s *ProductSuiteTest) SetupSuite() {
	pgPool, cleanUpfunc := db.ConnectTODBForSuite(config.Load())

	s.Repository = NewProductRepo(pgPool)
	s.ClenUpFunc = cleanUpfunc
}

func (s *ProductSuiteTest) TestProductCrud() {
	productCreate := pb.Product{
		Name:       "new product",
		Model:      "new model",
		TypeId:     1,
		CategoryId: 1,
		Price:      22.5,
		Amount:     5,
	}
	product, err := s.Repository.CreateProduct(&productCreate)
	s.Nil(err)
	s.NotNil(product)

	categoryCreate := pb.Category{
		Name: "newcategory",
	}

	category, err := s.Repository.CreateCategory(&categoryCreate)
	s.Nil(err)
	s.NotNil(category)

	typeCreate := pb.Type{
		Name: "newtype",
	}

	types, err := s.Repository.CreateType(&typeCreate)
	s.Nil(err)
	s.NotNil(types)

	getproductid := pb.Ids{Id: []int64{product.Id}}

	getproductbyid, err := s.Repository.GetProductInfoByid(&getproductid)
	s.Nil(err)
	s.NotNil(getproductbyid)

	updateIds := pb.Product{
		Id:   product.Id,
		Name: "updated name",
		// Model:      "new model",
		// TypeId:     1,
		// CategoryId: 1,
		// Price:      22.5,
		// Amount:     5,
	}
	updateId, err := s.Repository.UpdateByid(&updateIds)
	s.Nil(err)
	s.NotNil(updateId)
	// s.NotEqual(getproductbyid, updateId)

}

func (suite *ProductSuiteTest) TearDownSuite() {
	suite.ClenUpFunc()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductSuiteTest))
}
