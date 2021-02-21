package main

import (
	"net"

	"google.golang.org/grpc"
	// commonpb "github.com/preslavmihaylov/go-grpc-crash-course/gen/common"
	// "github.com/preslavmihaylov/go-grpc-crash-course/gen/payment_statements"
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
