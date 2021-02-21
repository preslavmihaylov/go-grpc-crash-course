package main

import (
	"errors"

	casinopb "github.com/preslavmihaylov/go-grpc-crash-course/gen/casino"

	// commonpb "github.com/preslavmihaylov/go-grpc-crash-course/gen/common"
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
