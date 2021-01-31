package main

import (
	"errors"
	"log"

	casinopb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/casino"

	// commonpb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/common"
	"google.golang.org/grpc"
)

type command string

const casinoAddr = "localhost:10000"

var errStopGambling = errors.New("user exits gambling session")

func setupClient() (casinopb.CasinoClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(casinoAddr, opts...)
	if err != nil {
		log.Fatalf("couldn't dial payment statements server: %v", err)
	}

	return casinopb.NewCasinoClient(conn), conn
}

func buyTokens(tokensCnt int) (string, error) {
	panic("not implemented")
}

func withdraw(tokensCnt int) (string, error) {
	panic("not implemented")
}

func tokenBalance() (string, error) {
	panic("not implemented")
}

func payments() (string, error) {
	panic("not implemented")
}

func paymentStatement() (string, error) {
	panic("not implemented")
}

func gamble() (string, error) {
	panic("not implemented")
}
