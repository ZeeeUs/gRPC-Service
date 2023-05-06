package grpc

import (
	"context"

	"github.com/rs/zerolog"

	pb "github.com/ZeeeUs/gRPC-Service/internal/domain/proto"
	"github.com/ZeeeUs/gRPC-Service/internal/domain/service"
	"github.com/ZeeeUs/gRPC-Service/internal/models"
)

type AutoMarket struct {
	pb.UnimplementedAutoMarketServer
	autoMarketService service.AutoMarketService
	log               zerolog.Logger
}

func (am *AutoMarket) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	newAccount := models.Account{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}

	id, err := am.autoMarketService.CreateAccount(ctx, newAccount)
	if err != nil {
		am.log.Error().Err(err).Msg("failed to create new account")
	}

	return &pb.CreateAccountResponse{Id: id}, nil
}

func New(log zerolog.Logger, autoMarketService service.AutoMarketService) *AutoMarket {
	return &AutoMarket{
		autoMarketService: autoMarketService,
		log:               log,
	}
}
