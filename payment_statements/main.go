package main

import (
	"net"

	"google.golang.org/grpc"
	// commonpb "gitlab.com/preslavmihaylov/go-comm-protocols-exercise/gen/common"
	// "gitlab.com/preslavmihaylov/go-comm-protocols-exercise/gen/payment_statements"
)

var (
	paymentStatementsAddr = "localhost:10001"
)

func main() {
	grpcServer, lis := setupPaymentStatementsServer()
	grpcServer.Serve(lis)
}

func setupPaymentStatementsServer() (*grpc.Server, net.Listener) {
	panic("not implemented")
}

type server struct{}
