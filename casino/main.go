package main

import (
	"context"
	"log"
	"net"

	casinopb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/casino"
	commonpb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/common"
	"gitlab.com/preslavmihaylov/go-grpc-exercise/gen/payment_statements"
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
	lis, err := net.Listen("tcp", casinoAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	casinopb.RegisterCasinoServer(grpcServer, newCasinoServer())

	return grpcServer, lis
}

func setupPaymentStatementsClient() (payment_statements.PaymentStatementsClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(paymentStatementsAddr, opts...)
	if err != nil {
		log.Fatalf("couldn't dial payment statements server: %v", err)
	}

	return payment_statements.NewPaymentStatementsClient(conn), conn
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

func (c *casinoServer) BuyTokens(_ context.Context, _ *commonpb.Payment) (*casinopb.Tokens, error) {
	panic("not implemented") // TODO: Implement
}

func (c *casinoServer) Withdraw(_ context.Context, _ *casinopb.WithdrawRequest) (*commonpb.Payment, error) {
	panic("not implemented") // TODO: Implement
}

func (c *casinoServer) GetTokenBalance(_ context.Context, _ *commonpb.User) (*casinopb.Tokens, error) {
	panic("not implemented") // TODO: Implement
}

func (c *casinoServer) GetPayments(_ *commonpb.User, _ casinopb.Casino_GetPaymentsServer) error {
	panic("not implemented") // TODO: Implement
}

func (c *casinoServer) GetPaymentStatement(_ context.Context, _ *commonpb.User) (*commonpb.PaymentStatement, error) {
	panic("not implemented") // TODO: Implement
}

func (c *casinoServer) Gamble(_ casinopb.Casino_GambleServer) error {
	panic("not implemented") // TODO: Implement
}
