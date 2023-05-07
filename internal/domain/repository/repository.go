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

	ctxDb, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	query := `INSERT INTO public.publication (user_id, brand, model, vin, production_year, mileage, pics_count, owner_count) 
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`
	err := am.conn.QueryRow(
		ctxDb,
		query,
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
			return 0, errors.Wrap(err, "insert new account query timeout")
		} else {
			return 0, err
		}
	}

	return id, nil
}

func New(log zerolog.Logger, conn *pgxpool.Pool) AutoMarketRepo {
	return &autoMarketRepo{
		conn: conn,
		log:  log,
	}
}
