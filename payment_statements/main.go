package main

import (
	"fmt"
	"log"
	"net"

	commonpb "github.com/preslavmihaylov/go-grpc-crash-course/gen/common"
	"github.com/preslavmihaylov/go-grpc-crash-course/gen/payment_statements"
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
	lis, err := net.Listen("tcp", paymentStatementsAddr)
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

func toPaymentStatement(payments []*commonpb.Payment) *commonpb.PaymentStatement {
	if len(payments) == 0 {
		return &commonpb.PaymentStatement{
			Data: "PAYMENT STATEMENT OF (unknown)\n\nNo earnings",
		}
	}

	statement := fmt.Sprintf("PAYMENT STATEMENT OF %s\n\n", payments[0].GetUser().GetId())
	statement += "PAYMENT HISTORY:\n"

	balance := 0
	for i, payment := range payments {
		statement += fmt.Sprintf("\tPayment %d: %d$\n", i, payment.Amount)
		balance += int(payment.Amount)
	}

	statement += fmt.Sprintf("\nFINAL BALANCE: %d\n", balance)

	return &commonpb.PaymentStatement{
		Data: statement,
	}
}
