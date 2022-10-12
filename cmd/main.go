package main

import (
	"github/FIrstService/template-service/Testproductservice/config"
	pb "github/FIrstService/template-service/Testproductservice/genproto/product"
	"github/FIrstService/template-service/Testproductservice/pkg/db"
	"github/FIrstService/template-service/Testproductservice/pkg/logger"
	"github/FIrstService/template-service/Testproductservice/service"

	grpcclient "github/FIrstService/template-service/Testproductservice/service/grpc_client"

	// grpcclient "github/FIrstService/template-service/Testproductservice/service/grpc_client"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "Testproductservice")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectTODB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	grpcClient, err := grpcclient.New(cfg)
	if err != nil {
		log.Fatal("error while connect to clients", logger.Error(err))
	}

	productService := service.NewProductService(grpcClient,connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterProductServiceServer(s, productService)
	log.Info("main: server runnig",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}
