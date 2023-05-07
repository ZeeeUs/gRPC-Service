package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/ZeeeUs/gRPC-Service/internal/models"
)

type AutoMarketRepo interface {
	CreatePublication(ctx context.Context, userID uint64, publication models.Publication) (uint64, error)
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

	publQuery := `INSERT INTO publication (brand, model, vin, production_year, mileage, pics_count, owner_count, description, is_active, created_at)
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

	charcQuery := `INSERT INTO car_characteristic (publication_id, body_type, gear_box, engine, engine_power, engine_capacity, drive_gear, color) 
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

func New(log zerolog.Logger, conn *pgxpool.Pool) AutoMarketRepo {
	return &autoMarketRepo{
		conn: conn,
		log:  log,
	}
}
