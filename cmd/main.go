package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ZeeeUs/gRPC-Service/internal/config"
	pb "github.com/ZeeeUs/gRPC-Service/internal/domain/proto"
	"github.com/ZeeeUs/gRPC-Service/internal/domain/repository"
	"github.com/ZeeeUs/gRPC-Service/internal/domain/service"
	transport "github.com/ZeeeUs/gRPC-Service/internal/domain/transport/grpc"
	"github.com/ZeeeUs/gRPC-Service/pkg/postgres"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg := config.GetConfig()
	pgConn, err := postgres.NewPoolConnection(
		context.Background(),
		cfg.DatabaseConfig.Name,
		cfg.DatabaseConfig.Address,
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.MaxIdleLifetime,
		cfg.DatabaseConfig.MaxLifetime,
		cfg.DatabaseConfig.PrepareCacheCap,
		cfg.DatabaseConfig.MaxConn,
	)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: cfg.GRPCConfig.MaxConnectionIdle,
		Timeout:           cfg.GRPCConfig.Timeout,
		MaxConnectionAge:  cfg.GRPCConfig.MaxConnectionAge,
	}))

	repo := repository.New(log.Logger, pgConn)
	svc := service.New(log.Logger, repo)
	server := transport.New(log.Logger, svc)

	listener, err := net.Listen(cfg.GRPCConfig.Network, cfg.GRPCConfig.Address)
	if err != nil {
		log.Error().Err(err).Msg("can't start listener")
	}
	pb.RegisterAutoMarketServer(grpcServer, server)

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			log.Error().Err(err).Msg("failed to serve gRPC server")
			cancel()
		}
	}()

	go func() {
		mux := runtime.NewServeMux()
		if err := pb.RegisterAutoMarketHandlerServer(context.Background(), mux, server); err != nil {
			log.Error().Err(err).Msg("failed to register http server")
			cancel()
		}

		if err = http.ListenAndServe("localhost:8080", mux); err != nil {
			log.Error().Err(err).Msg("failed to start http server")
			cancel()
		}
	}()

	select {
	case <-shutdown:
		log.Info().Msgf("start shutdown server by sys call")
	case <-ctx.Done():
		log.Error().Msgf("start shutdown server by context")
	}

	grpcServer.GracefulStop()
}
