package storage

import (
	"github/FIrstService/template-service/Testproductservice/storage/postgres"
	"github/FIrstService/template-service/Testproductservice/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Product() repo.ProductStorageI
}

type storagePg struct {
	db       *sqlx.DB
	productRepo repo.ProductStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		productRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
