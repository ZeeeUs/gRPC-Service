package service

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/ZeeeUs/gRPC-Service/internal/domain/repository"
	"github.com/ZeeeUs/gRPC-Service/internal/models"
)

type AutoMarketService interface {
	CreatePublication(ctx context.Context, userID uint64, newPublication models.Publication) (uint64, error)
	GetColors(ctx context.Context) ([]models.Color, error)
}

type autoMarketService struct {
	autoMarketRepo repository.AutoMarketRepo
	log            zerolog.Logger
}

func (am *autoMarketService) CreatePublication(ctx context.Context, userID uint64, publication models.Publication) (uint64, error) {
	return am.autoMarketRepo.CreatePublication(ctx, userID, publication)
}

func (am *autoMarketService) GetColors(ctx context.Context) ([]models.Color, error) {
	return am.autoMarketRepo.GetColors(ctx)
}

func New(log zerolog.Logger, autoMarketRepo repository.AutoMarketRepo) AutoMarketService {
	return &autoMarketService{
		autoMarketRepo: autoMarketRepo,
		log:            log,
	}
}
