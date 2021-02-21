package main

import (
	"errors"

	"google.golang.org/grpc"
	// casinopb "github.com/preslavmihaylov/go-grpc-crash-course/gen/casino"
	// commonpb "github.com/preslavmihaylov/go-grpc-crash-course/gen/common"
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
