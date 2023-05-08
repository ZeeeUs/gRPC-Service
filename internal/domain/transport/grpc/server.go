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

func (am *AutoMarket) GetEngines(ctx context.Context, _ *pb.GetEnginesRequest) (*pb.GetEnginesResponse, error) {
	engines, err := am.autoMarketService.GetEngines(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*pb.Engine, 0)
	for _, engine := range engines {
		resp = append(resp, &pb.Engine{
			Id:   engine.ID,
			Name: engine.Name,
		})
	}

	return &pb.GetEnginesResponse{Engines: resp}, nil
}

func (am *AutoMarket) GetGearBoxes(ctx context.Context, _ *pb.GetGearBoxesRequest) (*pb.GetGearBoxesResponse, error) {
	gearBoxes, err := am.autoMarketService.GetGearBoxes(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*pb.GearBox, 0)
	for _, gearBox := range gearBoxes {
		resp = append(resp, &pb.GearBox{
			Id:   gearBox.ID,
			Name: gearBox.Name,
		})
	}

	return &pb.GetGearBoxesResponse{GearBoxes: resp}, nil
}

func (am *AutoMarket) GetBodyTypes(ctx context.Context, _ *pb.GetBodyTypesRequest) (*pb.GetBodyTypesResponse, error) {
	bodyTypes, err := am.autoMarketService.GetBodyTypes(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*pb.BodyType, 0)
	for _, bodyType := range bodyTypes {
		resp = append(resp, &pb.BodyType{
			Id:   bodyType.ID,
			Name: bodyType.Name,
		})
	}

	return &pb.GetBodyTypesResponse{BodyTypes: resp}, nil
}

func (am *AutoMarket) GetBrands(ctx context.Context, _ *pb.GetBrandsRequest) (*pb.GetBrandsResponse, error) {
	brands, err := am.autoMarketService.GetBrands(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*pb.Brand, 0)
	for _, brand := range brands {
		resp = append(resp, &pb.Brand{
			Id:   brand.ID,
			Name: brand.Name,
		})
	}

	return &pb.GetBrandsResponse{Brands: resp}, nil
}

func (am *AutoMarket) GetDriveGears(ctx context.Context, _ *pb.GetDriveGearsRequest) (*pb.GetDriveGearsResponse, error) {
	driveGears, err := am.autoMarketService.GetDriveGears(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*pb.DriveGear, 0)
	for _, driveGear := range driveGears {
		resp = append(resp, &pb.DriveGear{
			Id:   driveGear.ID,
			Name: driveGear.Name,
		})
	}

	return &pb.GetDriveGearsResponse{DriveGears: resp}, nil
}

func (am *AutoMarket) GetModels(ctx context.Context, _ *pb.GetModelsRequest) (*pb.GetModelsResponse, error) {
	carModels, err := am.autoMarketService.GetModels(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*pb.Model, 0)
	for _, model := range carModels {
		resp = append(resp, &pb.Model{
			Id:       model.ID,
			Name:     model.Name,
			BrandId:  model.BrandID,
			ParentId: model.ParentID,
		})
	}

	return &pb.GetModelsResponse{Models: resp}, nil
}

func New(log zerolog.Logger, autoMarketService service.AutoMarketService) *AutoMarket {
	return &AutoMarket{
		autoMarketService: autoMarketService,
		log:               log,
	}
}
