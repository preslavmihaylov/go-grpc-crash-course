package main

import (
	"log"
	"net"

	"gitlab.com/preslavmihaylov/go-grpc-exercise/gen/payment_statements"
	// commonpb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/common"
	"google.golang.org/grpc"
)

var (
	paymentStatementsAddr = "localhost:10001"
)

func main() {
	grpcServer, lis := setupPaymentStatementsServer()
	grpcServer.Serve(lis)
}

func setupPaymentStatementsServer() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", "localhost:10001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	payment_statements.RegisterPaymentStatementsServer(grpcServer, &server{})

	return grpcServer, lis
}

type server struct{}

func (s *server) CreateStatement(stream payment_statements.PaymentStatements_CreateStatementServer) error {
	panic("not implemented")
}
