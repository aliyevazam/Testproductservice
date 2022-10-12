package postgres

import (
	"fmt"
	pb "github/FIrstService/template-service/Testproductservice/genproto/product"
	"log"

	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

// NewProductRepo ...

func NewProductRepo(db *sqlx.DB) *productRepo {
	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(product *pb.ProductFullInfo) (*pb.ProductFullInfoResponse, error) {
	productResp := pb.ProductFullInfoResponse{}
	err := r.db.QueryRow(`insert into products (name,model,typeid,categoryid,price,amount) values ($1,$2,$3,$4,$5,$6)returning id,name,model,typeid,categoryid,price,amount`, product.Name, product.Model, product.TypeId, product.CategoryId, product.Price, product.Amount).Scan(
		&productResp.Id, &productResp.Name, &product.Model, &productResp.TypeId, &productResp.CategoryId, &productResp.Price, &productResp.Amount)
	if err != nil {
		return &pb.ProductFullInfoResponse{}, err
	}
	return &productResp, nil
}

func (r *productRepo) CreateCategory(req *pb.Category) (*pb.Category, error) {
	query := pb.Category{}
	err := r.db.QueryRow(`insert into categories (name) values($1)returning name`, req.Name).Scan(&req.Name)
	if err != nil {
		return &pb.Category{}, err
	}
	return &query, nil
}

func (r *productRepo) CreateType(req *pb.Type) (*pb.Type, error) {
	query := pb.Type{}
	err := r.db.QueryRow(`insert into types (name) values($1)returning name`, req.Name).Scan(&req.Name)
	if err != nil {
		return &pb.Type{}, err
	}
	return &query, nil
}

func (r *productRepo) GetProductInfoByid(ids *pb.Ids) (*pb.GetProducts, error) {
	fmt.Println(ids)
	response := &pb.GetProducts{}
	for _, id := range ids.Id {
		tempUser := &pb.Product{}
		err := r.db.QueryRow(`select * from products where id=$1`, id).Scan(&tempUser.Id, &tempUser.Name, &tempUser.Model, &tempUser.TypeId, &tempUser.CategoryId, &tempUser.Price, &tempUser.Amount)
		if err != nil {
			log.Fatal("Error while select products", err)
		}
		response.Products = append(response.Products, tempUser)
	}
	return response, nil
}

func (r *productRepo) UpdateByid(req *pb.Product) (*pb.Product, error) {
	_, err := r.db.Exec(`UPDATE products SET name=$1, model=$2 where id=$3`,
		req.Name, req.Model, req.Id)
	fmt.Println(err)
	return req, err
}

func (r *productRepo) DeleteInfo(ids *pb.Ids) error {
	for _, id := range ids.Id {
		_, err := r.db.Exec(`DELETE FROM products WHERE id=$1`, id)
		if err != nil {
			log.Fatal("Error while delete product", err)
		}
	}
	return nil
}
