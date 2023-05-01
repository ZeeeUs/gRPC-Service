package main

import (
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/ZeeeUs/gRPC-Service/internal/social_network/proto"
	"github.com/ZeeeUs/gRPC-Service/internal/social_network/repository"
	socialNetwork "github.com/ZeeeUs/gRPC-Service/internal/social_network/transport/grpc"
	"github.com/ZeeeUs/gRPC-Service/internal/social_network/usecase"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Error().Err(err).Msg("can't start listener")
	}

	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 15 * time.Minute,
		Timeout:           30 * time.Second,
		MaxConnectionAge:  15 * time.Minute,
	}))

	repo := repository.New(log.Logger)
	uc := usecase.New(log.Logger, repo)
	svc := socialNetwork.New(log.Logger, uc)

	pb.RegisterSocialNetworkServer(grpcServer, svc)

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			log.Error().Err(err).Msg("failed to serve gRPC server")
		}
	}()
}
