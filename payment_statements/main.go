package main

import (
	"fmt"
	"log"
	"net"

	// commonpb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/common"
	"gitlab.com/preslavmihaylov/go-grpc-exercise/gen/payment_statements"
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

func setupPaymentStatementsServer() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", fmt.Sprintf(paymentStatementsAddr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	payment_statements.RegisterPaymentStatementsServer(grpcServer, &server{})

	return grpcServer, lis
}

type server struct{}

func (s *server) CreateStatement(stream payment_statements.PaymentStatements_CreateStatementServer) error {
	panic("not implemented") // TODO: Implement
}
