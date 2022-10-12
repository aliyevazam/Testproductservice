package grpcClient

import (
	"fmt"
	"github/FIrstService/template-service/Testproductservice/config"
	storePB "github/FIrstService/template-service/Testproductservice/genproto/stores"

	"google.golang.org/grpc"
)

type ServiceManager struct {
	conf         config.Config
	storeservice storePB.StoreServiceClient
}

// // GrpcClientI ...
// type GrpcClientI interface {
// }

// GrpcClient ...
// type GrpcClient struct {
// 	cfg         config.Config
// 	connections map[string]interface{}
// }

// New ...
// func New(cfg config.Config) (*GrpcClient, error) {
// 	return &GrpcClient{
// 		cfg:         cfg,
// 		connections: map[string]interface{}{},
// 	}, nil
// }

func New(cnfg config.Config) (*ServiceManager, error) {
	connStore, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cnfg.StoreServiceHost, cnfg.StoreServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("error while dial product service: host: %s and port: %d",
			cnfg.ReviewServiceHost, cnfg.ReviewServicePort)

	}
	serviceManager := &ServiceManager{
		conf:         cnfg,
		storeservice: storePB.NewStoreServiceClient(connStore),
	}

	return serviceManager, nil
}

func (s *ServiceManager) StoreService() storePB.StoreServiceClient {
	return s.storeservice
}
