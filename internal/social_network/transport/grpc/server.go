package grpc

import (
	"context"

	"google.golang.org/grpc/status"

	"github.com/ZeeeUs/gRPC-Service/internal/models"
	pb "github.com/ZeeeUs/gRPC-Service/internal/social_network/proto"
	"github.com/ZeeeUs/gRPC-Service/internal/social_network/service"
	"github.com/ZeeeUs/gRPC-Service/pkg/grpc_errors"
	"github.com/rs/zerolog"
)

type SocialNetwork struct {
	pb.UnimplementedSocialNetworkServer
	socialNetworkService service.SocialNetworkService
	log                  zerolog.Logger
}

func (sn *SocialNetwork) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	newAccount := models.Account{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	id, err := sn.socialNetworkService.CreateAccount(ctx, newAccount)
	if err != nil {
		sn.log.Error().Err(err).Msg("failed to create new account")
		return nil, status.Error(grpc_errors.ParseGRPCErrStatusCode(err), "invalid request")
	}
	return &pb.CreateAccountResponse{Id: id}, nil
}

func New(log zerolog.Logger, socialNetworkService service.SocialNetworkService) *SocialNetwork {
	return &SocialNetwork{
		socialNetworkService: socialNetworkService,
		log:                  log,
	}
}
