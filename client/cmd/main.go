package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	socialNetwork "github.com/ZeeeUs/gRPC-Service/social_network/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connetc to server: %s", err)
	}

	client := socialNetwork.NewSocialNetworkClient(conn)

	var newUser = socialNetwork.CreateAccountRequest{
		Name:  "John",
		Email: "john@test.go",
		Age:   28,
	}
	resp, err := client.CreateAccount(context.Background(), &newUser)
	if err != nil {
		log.Printf("failed to create account: %s", err)
	}

	fmt.Println(resp.Id)
}
