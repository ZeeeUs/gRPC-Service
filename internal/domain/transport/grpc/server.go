package grpc

import (
	"context"
	"time"

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

func (am *AutoMarket) CreatePublication(ctx context.Context, req *pb.CreatePublicationRequest) (*pb.CreatePublicationResponse, error) {
	layout := "02-01-2006"
	date, err := time.Parse(layout, req.ProductionYear)
	if err != nil {
		am.log.Error().Err(err).Msg("failed to convert string to date")
		return nil, err
	}

	newPublication := models.Publication{
		Brand:          req.Brand,
		Model:          req.Model,
		Vin:            req.Vin,
		ProductionYear: date,
		Mileage:        req.Mileage,
		PicsCount:      req.PicsCount,
		OwnerCount:     req.OwnerCount,
		Color:          req.Color,
		BodyType:       req.BodyType,
		DriveGear:      req.DriveGear,
		GearBox:        req.GearBox,
		EngineType:     req.EngineType,
		EngineCapacity: req.EngineCapacity,
		EnginePower:    req.EnginePower,
		Description:    req.Description,
	}

	// TODO добавить поддержку userID
	userID := uint64(55)

	id, err := am.autoMarketService.CreatePublication(ctx, userID, newPublication)
	if err != nil {
		am.log.Error().Err(err).Msg("failed to create new account")
		return nil, err
	}

	return &pb.CreatePublicationResponse{Id: id}, nil
}

func (am *AutoMarket) GetColors(ctx context.Context, _ *pb.GetColorsRequest) (*pb.GetColorsResponse, error) {
	colors, err := am.autoMarketService.GetColors(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*pb.Color, 0)
	for _, color := range colors {
		resp = append(resp, &pb.Color{
			Id:      color.ID,
			Name:    color.Name,
			HexCode: color.HexCode,
		})
	}

	return &pb.GetColorsResponse{Colors: resp}, nil
}

func New(log zerolog.Logger, autoMarketService service.AutoMarketService) *AutoMarket {
	return &AutoMarket{
		autoMarketService: autoMarketService,
		log:               log,
	}
}
