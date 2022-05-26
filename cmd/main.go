package main

import (
	"net"
	"google.golang.org/grpc/reflection"

	"github.com/baxromumarov/template-service/config"
	pb "github.com/baxromumarov/template-service/genproto"
	"github.com/baxromumarov/template-service/pkg/db"
	"github.com/baxromumarov/template-service/pkg/logger"
	"github.com/baxromumarov/template-service/service"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := service.NewUserService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}

// reflection learning