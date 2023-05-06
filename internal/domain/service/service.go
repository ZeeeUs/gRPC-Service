package service

import (
	"context"

	"github.com/ZeeeUs/gRPC-Service/internal/domain/repository"
	"github.com/ZeeeUs/gRPC-Service/internal/models"
	"github.com/rs/zerolog"
)

type AutoMarketService interface {
	CreateAccount(ctx context.Context, account models.Account) (int64, error)
}

type autoMarketService struct {
	autoMarketRepo repository.AutoMarketRepo
	log            zerolog.Logger
}

func (am *autoMarketService) CreateAccount(ctx context.Context, account models.Account) (int64, error) {
	return am.autoMarketRepo.CreateAccount(ctx, account)
}

func New(log zerolog.Logger, autoMarketRepo repository.AutoMarketRepo) AutoMarketService {
	return &autoMarketService{
		autoMarketRepo: autoMarketRepo,
		log:            log,
	}
}
