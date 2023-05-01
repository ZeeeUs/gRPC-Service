package service

import (
	"context"

	"github.com/ZeeeUs/gRPC-Service/internal/domain/repository"
	"github.com/ZeeeUs/gRPC-Service/internal/models"
	"github.com/rs/zerolog"
)

type SocialNetworkService interface {
	CreateAccount(ctx context.Context, account models.Account) (int64, error)
}

type socialNetworkService struct {
	socialNetworkRepo repository.SocialNetworkRepo
	log               zerolog.Logger
}

func (snu *socialNetworkService) CreateAccount(ctx context.Context, account models.Account) (int64, error) {
	return snu.socialNetworkRepo.CreateAccount(ctx, account)
}

func New(log zerolog.Logger, repo repository.SocialNetworkRepo) SocialNetworkService {
	return &socialNetworkService{
		socialNetworkRepo: repo,
		log:               log,
	}
}
