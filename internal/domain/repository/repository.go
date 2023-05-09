package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/ZeeeUs/gRPC-Service/internal/models"
)

type AutoMarketRepo interface {
	CreatePublication(ctx context.Context, userID uint64, publication models.Publication) (uint64, error)
	GetColors(ctx context.Context) ([]models.Color, error)
	GetEngines(ctx context.Context) ([]models.Engine, error)
	GetGearBoxes(ctx context.Context) ([]models.GearBox, error)
	GetBodyTypes(ctx context.Context) ([]models.BodyType, error)
	GetBrands(ctx context.Context) ([]models.Brand, error)
	GetDriveGears(ctx context.Context) ([]models.DriveGear, error)
	GetModels(ctx context.Context, brandId uint64) ([]models.Model, error)
}

type autoMarketRepo struct {
	conn *pgxpool.Pool
	log  zerolog.Logger
}

func (am *autoMarketRepo) CreatePublication(ctx context.Context, userID uint64, publication models.Publication) (uint64, error) {
	var id uint64

	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	tx, err := am.conn.Begin(ctx)
	if err != nil {
		am.log.Error().Err(err).Msg("failed to create transaction for create publication")
		return 0, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	publQuery := `INSERT INTO publications (brand, model, vin, production_year, mileage, pics_count, owner_count, description, is_active, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,true,now()) RETURNING id`
	err = tx.QueryRow(
		ctxDb,
		publQuery,
		userID,
		publication.Brand,
		publication.Model,
		publication.Vin,
		publication.ProductionYear,
		publication.Mileage,
		publication.PicsCount,
		publication.OwnerCount,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return 0, errors.Wrap(err, "insert new publication query timeout")
		}
		return 0, err
	}

	charcQuery := `INSERT INTO car_characteristics (publication_id, body_type, gear_box, engine, engine_power, engine_capacity, drive_gear, color) 
					VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err = tx.Exec(
		ctxDb,
		charcQuery,
		id,
		publication.BodyType,
		publication.GearBox,
		publication.EngineType,
		publication.EngineCapacity,
		publication.DriveGear,
		publication.Color,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return 0, errors.Wrap(err, "insert new publication characteristics query timeout")
		}
		return 0, err
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, errors.Wrap(err, "failed to commit when add new publication")
	}

	return id, nil
}

func (am *autoMarketRepo) GetColors(ctx context.Context) ([]models.Color, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	query := "SELECT id, name, hex_code FROM public.colors"
	rows, err := am.conn.Query(ctxDb, query)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, errors.Wrap(err, "get colors query timeout")
		}
		return nil, err
	}
	defer rows.Close()

	var colors = make([]models.Color, 0)
	for rows.Next() {
		var color models.Color
		if err = rows.Scan(&color.ID, &color.Name, &color.HexCode); err != nil {
			return nil, errors.Wrap(err, "failed to scan colors")
		}
		colors = append(colors, color)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return colors, nil
}

func (am *autoMarketRepo) GetEngines(ctx context.Context) ([]models.Engine, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	query := "SELECT id, name FROM public.engines"
	rows, err := am.conn.Query(ctxDb, query)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, errors.Wrap(err, "get engines query timeout")
		}
		return nil, err
	}
	defer rows.Close()

	var engines = make([]models.Engine, 0)
	for rows.Next() {
		var engine models.Engine
		if err = rows.Scan(&engine.ID, &engine.Name); err != nil {
			return nil, errors.Wrap(err, "failed to scan engines")
		}
		engines = append(engines, engine)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return engines, nil
}

func (am *autoMarketRepo) GetGearBoxes(ctx context.Context) ([]models.GearBox, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	query := "SELECT id, name FROM public.gear_boxes"
	rows, err := am.conn.Query(ctxDb, query)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, errors.Wrap(err, "get gear boxes query timeout")
		}
		return nil, err
	}
	defer rows.Close()

	var gerBoxes = make([]models.GearBox, 0)
	for rows.Next() {
		var gearBox models.GearBox
		if err = rows.Scan(&gearBox.ID, &gearBox.Name); err != nil {
			return nil, errors.Wrap(err, "failed to scan gear boxes")
		}
		gerBoxes = append(gerBoxes, gearBox)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gerBoxes, nil
}

func (am *autoMarketRepo) GetBodyTypes(ctx context.Context) ([]models.BodyType, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	query := "SELECT id, name FROM public.body_types"
	rows, err := am.conn.Query(ctxDb, query)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, errors.Wrap(err, "get body types query timeout")
		}
		return nil, err
	}
	defer rows.Close()

	var bodyTypes = make([]models.BodyType, 0)
	for rows.Next() {
		var bodyType models.BodyType
		if err = rows.Scan(&bodyType.ID, &bodyType.Name); err != nil {
			return nil, errors.Wrap(err, "failed to body types")
		}
		bodyTypes = append(bodyTypes, bodyType)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bodyTypes, nil
}

func (am *autoMarketRepo) GetBrands(ctx context.Context) ([]models.Brand, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	query := "SELECT id, name FROM public.brands"
	rows, err := am.conn.Query(ctxDb, query)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, errors.Wrap(err, "get brands query timeout")
		}
		return nil, err
	}
	defer rows.Close()

	var brands = make([]models.Brand, 0)
	for rows.Next() {
		var brand models.Brand
		if err = rows.Scan(&brand.ID, &brand.Name); err != nil {
			return nil, errors.Wrap(err, "failed to scan brands")
		}
		brands = append(brands, brand)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return brands, nil
}

func (am *autoMarketRepo) GetDriveGears(ctx context.Context) ([]models.DriveGear, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	query := "SELECT id, name FROM public.drive_gears"
	rows, err := am.conn.Query(ctxDb, query)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, errors.Wrap(err, "get drive gears query timeout")
		}
		return nil, err
	}
	defer rows.Close()

	var driveGears = make([]models.DriveGear, 0)
	for rows.Next() {
		var driveGear models.DriveGear
		if err = rows.Scan(&driveGear.ID, &driveGear.Name); err != nil {
			return nil, errors.Wrap(err, "failed to scan drive gears")
		}
		driveGears = append(driveGears, driveGear)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return driveGears, nil
}

func (am *autoMarketRepo) GetModels(ctx context.Context, brandId uint64) ([]models.Model, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	query := "SELECT id, name, brand_id, parent_id FROM public.models"
	if brandId > 0 {
		query += fmt.Sprintf(" WHERE brand_id=%d", brandId)
	}

	rows, err := am.conn.Query(ctxDb, query)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, errors.Wrap(err, "get models query timeout")
		}
		return nil, err
	}
	defer rows.Close()

	var carModels = make([]models.Model, 0)
	for rows.Next() {
		var carModel models.Model
		if err = rows.Scan(&carModel.ID, &carModel.Name, &carModel.BrandID, &carModel.ParentID); err != nil {
			return nil, errors.Wrap(err, "failed to scan models")
		}
		carModels = append(carModels, carModel)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return carModels, nil
}

func New(log zerolog.Logger, conn *pgxpool.Pool) AutoMarketRepo {
	return &autoMarketRepo{
		conn: conn,
		log:  log,
	}
}
