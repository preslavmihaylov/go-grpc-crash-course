package main

import (
	"errors"

	"google.golang.org/grpc"
	// casinopb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/casino"
	// commonpb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/common"
)

type command string

var errStopGambling = errors.New("user exits gambling session")

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
