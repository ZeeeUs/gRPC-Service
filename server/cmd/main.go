package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	socialNetwork "github.com/ZeeeUs/gRPC-Service/social_network/proto"
)

type SocialNetworkServer struct {
	socialNetwork.UnimplementedSocialNetworkServer
}

func (sn *SocialNetworkServer) CreateAccount(context.Context, *socialNetwork.CreateAccountRequest) (*socialNetwork.CreateAccountResponse, error) {
	return &socialNetwork.CreateAccountResponse{Id: 23}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("can't start listener: %s", err)
	}

	grpcServer := grpc.NewServer()

	socialNetwork.RegisterSocialNetworkServer(grpcServer, &SocialNetworkServer{})
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server grpcServer: %s", err)
	}
}
