package main

import (
	"log"
	"net"

	// casinopb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/casino"
	// commonpb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/common"
	// "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/payment_statements"
	"google.golang.org/grpc"
)

type userID string

var (
	tokensPerDollar         = int32(5)
	casinoAddr              = "localhost:10000"
	paymentStatementsAddr   = "localhost:10001"
	paymentStatementsClient payment_statements.PaymentStatementsClient
)

func main() {
	var conn *grpc.ClientConn
	paymentStatementsClient, conn = setupPaymentStatementsClient()
	defer conn.Close()

	log.Println("Successfully connected to payment_statements...")

	grpcServer, lis := setupCasinoServer()
	grpcServer.Serve(lis)
}

func setupCasinoServer() (*grpc.Server, net.Listener) {
	panic("not implemented")
}

func setupPaymentStatementsClient() (payment_statements.PaymentStatementsClient, *grpc.ClientConn) {
	panic("not implemented")
}

func newCasinoServer() *casinoServer {
	return &casinoServer{
		stockPrice:     10,
		userToTokens:   map[userID]int32{},
		userToPayments: map[userID][]int32{},
		userToStocks:   map[userID]int32{},
	}
}

type casinoServer struct {
	stockPrice int32

	userToTokens   map[userID]int32
	userToPayments map[userID][]int32
	userToStocks   map[userID]int32
}
