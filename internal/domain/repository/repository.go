package repository

import (
	"context"

	"github.com/ZeeeUs/gRPC-Service/internal/models"
	"github.com/rs/zerolog"
)

type AutoMarketRepo interface {
	CreateAccount(ctx context.Context, account models.Account) (int64, error)
}

type autoMarketRepo struct {
	log zerolog.Logger
}

func (am *autoMarketRepo) CreateAccount(ctx context.Context, account models.Account) (int64, error) {
	return 23, nil
}

func New(log zerolog.Logger) AutoMarketRepo {
	return &autoMarketRepo{
		log: log,
	}
}
