package repository

import (
	"context"

	"github.com/ZeeeUs/gRPC-Service/internal/models"
	"github.com/rs/zerolog"
)

type SocialNetworkRepo interface {
	CreateAccount(ctx context.Context, account models.Account) (int64, error)
}

type socialNetworkRepo struct {
	log zerolog.Logger
}

func (snr *socialNetworkRepo) CreateAccount(ctx context.Context, account models.Account) (int64, error) {
	//query := "INSERT INTO domain.account (name, email, age) VALUES ($1, $2, $3)"
	return 23, nil
}

func New(log zerolog.Logger) SocialNetworkRepo {
	return &socialNetworkRepo{
		log: log,
	}
}
