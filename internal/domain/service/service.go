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
	GetEngines(ctx context.Context) ([]models.Engine, error)
	GetGearBoxes(ctx context.Context) ([]models.GearBox, error)
	GetBodyTypes(ctx context.Context) ([]models.BodyType, error)
	GetBrands(ctx context.Context) ([]models.Brand, error)
	GetDriveGears(ctx context.Context) ([]models.DriveGear, error)
	GetModels(ctx context.Context, brandId uint64) ([]models.Model, error)
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

func (am *autoMarketService) GetEngines(ctx context.Context) ([]models.Engine, error) {
	return am.autoMarketRepo.GetEngines(ctx)
}

func (am *autoMarketService) GetGearBoxes(ctx context.Context) ([]models.GearBox, error) {
	return am.autoMarketRepo.GetGearBoxes(ctx)
}

func (am *autoMarketService) GetBodyTypes(ctx context.Context) ([]models.BodyType, error) {
	return am.autoMarketRepo.GetBodyTypes(ctx)
}

func (am *autoMarketService) GetBrands(ctx context.Context) ([]models.Brand, error) {
	return am.autoMarketRepo.GetBrands(ctx)
}

func (am *autoMarketService) GetDriveGears(ctx context.Context) ([]models.DriveGear, error) {
	return am.autoMarketRepo.GetDriveGears(ctx)
}

func (am *autoMarketService) GetModels(ctx context.Context, brandId uint64) ([]models.Model, error) {
	return am.autoMarketRepo.GetModels(ctx, brandId)
}

func New(log zerolog.Logger, autoMarketRepo repository.AutoMarketRepo) AutoMarketService {
	return &autoMarketService{
		autoMarketRepo: autoMarketRepo,
		log:            log,
	}
}
