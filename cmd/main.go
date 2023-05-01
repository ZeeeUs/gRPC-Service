package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/ZeeeUs/gRPC-Service/internal/config"
	pb "github.com/ZeeeUs/gRPC-Service/internal/social_network/proto"
	"github.com/ZeeeUs/gRPC-Service/internal/social_network/repository"
	"github.com/ZeeeUs/gRPC-Service/internal/social_network/service"
	grpcTransport "github.com/ZeeeUs/gRPC-Service/internal/social_network/transport/grpc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Msgf("failed to get config: %v", err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: cfg.GRPCConfig.MaxConnectionIdle,
		Timeout:           cfg.GRPCConfig.Timeout,
		MaxConnectionAge:  cfg.GRPCConfig.MaxConnectionAge,
	}))

	repo := repository.New(log.Logger)
	svc := service.New(log.Logger, repo)
	server := grpcTransport.New(log.Logger, svc)

	listener, err := net.Listen(cfg.GRPCConfig.Network, cfg.GRPCConfig.Address)
	if err != nil {
		log.Error().Err(err).Msg("can't start listener")
	}
	pb.RegisterSocialNetworkServer(grpcServer, server)

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			log.Error().Err(err).Msg("failed to serve gRPC server")
			cancel()
		}
	}()

	select {
	case v := <-shutdown:
		log.Info().Msgf("start shutdown server: %v", v)
	case done := <-ctx.Done():
		log.Error().Msgf("ctx.Done: %v", done)
	}

	grpcServer.GracefulStop()
}
