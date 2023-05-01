package usecase

import (
	"context"

	"github.com/ZeeeUs/gRPC-Service/internal/models"
	"github.com/ZeeeUs/gRPC-Service/internal/social_network/repository"
	"github.com/rs/zerolog"
)

type SocialNetworkUsecase interface {
	CreateAccount(ctx context.Context, account models.Account) (int64, error)
}

type socialNetworkUsecase struct {
	socialNetworkRepo repository.SocialNetworkRepo
	log               zerolog.Logger
}

func (snu *socialNetworkUsecase) CreateAccount(ctx context.Context, account models.Account) (int64, error) {
	return snu.socialNetworkRepo.CreateAccount(ctx, account)
}

func New(log zerolog.Logger, repo repository.SocialNetworkRepo) SocialNetworkUsecase {
	return &socialNetworkUsecase{
		socialNetworkRepo: repo,
		log:               log,
	}
}
