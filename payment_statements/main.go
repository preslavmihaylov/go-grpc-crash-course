package main

import (
	"log"
	"net"

	// commonpb "github.com/preslavmihaylov/go-grpc-crash-course/gen/common"
	// "github.com/preslavmihaylov/go-grpc-crash-course/gen/payment_statements"
	"google.golang.org/grpc"
)

var (
	paymentStatementsAddr = "localhost:10001"
)

func main() {
	grpcServer, lis := setupPaymentStatementsServer()

	log.Println("Successfully started payment_statements grpc server...")
	grpcServer.Serve(lis)
}

// TODO: Setup the payment statements grpc server
func setupPaymentStatementsServer() (*grpc.Server, net.Listener) {
	panic("not implemented")
}

// TODO: Implement the payment_statements.PaymentStatementsServer interface
type server struct{}
