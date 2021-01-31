package main

import (
	"errors"

	casinopb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/casino"

	// commonpb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/common"
	"google.golang.org/grpc"
)

type command string

const casinoAddr = "localhost:10000"

var errStopGambling = errors.New("user exits gambling session")

// TODO: Setup the casinopb.CasinoClient
func setupClient() (casinopb.CasinoClient, *grpc.ClientConn) {
	panic("not implemented")
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
