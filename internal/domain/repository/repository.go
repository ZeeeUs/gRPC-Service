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
	CreateAccount(ctx context.Context, account models.Account) (int64, error)
}

type autoMarketRepo struct {
	conn *pgxpool.Pool
	log  zerolog.Logger
}

func (am *autoMarketRepo) CreateAccount(ctx context.Context, account models.Account) (int64, error) {
	var id int64

	ctxDb, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	query := "INSERT INTO public.users (name, age, email) VALUES ($1, $2, $3) RETURNING id"
	err := am.conn.QueryRow(ctxDb, query, account.Name, account.Age, account.Email).Scan(&id)
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
