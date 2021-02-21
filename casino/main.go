package main

import (
	"context"
	"fmt"
	"log"
	"net"

	casinopb "github.com/preslavmihaylov/go-grpc-crash-course/gen/casino"
	commonpb "github.com/preslavmihaylov/go-grpc-crash-course/gen/common"
	"github.com/preslavmihaylov/go-grpc-crash-course/gen/payment_statements"
	"google.golang.org/grpc"
)

type userID string
type streamHandler func(stream casinopb.Casino_GambleServer) error

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
		log.Fatalf("fail to dial: %v", err)
	}

	return payment_statements.NewPaymentStatementsClient(conn), conn
}

func newCasinoServer() *casinoServer {
	return &casinoServer{
		stockPrice:     100,
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

func (c *casinoServer) BuyTokens(ctx context.Context, payment *commonpb.Payment) (*casinopb.Tokens, error) {
	log.Printf("BuyTokens invoked with payment %v\n", payment)

	usrID := userID(payment.User.GetId())
	tokens := payment.GetAmount() * tokensPerDollar

	c.userToTokens[usrID] += tokens
	c.userToPayments[usrID] = append(c.userToPayments[usrID], -payment.Amount)

	return &casinopb.Tokens{Count: tokens}, nil
}

func (c *casinoServer) Withdraw(ctx context.Context, withdrawReq *casinopb.WithdrawRequest) (*commonpb.Payment, error) {
	toWithdraw := withdrawReq.GetTokensCnt()
	log.Printf("Withdraw invoked with tokens %v\n", toWithdraw)

	usrID := userID(withdrawReq.User.GetId())
	log.Println(c.userToTokens[usrID])
	if !c.hasEnoughTokens(usrID, toWithdraw) {
		return nil, fmt.Errorf("not enough tokens to withdraw")
	}

	amount := toWithdraw / tokensPerDollar
	c.userToTokens[usrID] -= toWithdraw
	c.userToPayments[usrID] = append(c.userToPayments[usrID], amount)

	return &commonpb.Payment{User: withdrawReq.User, Amount: amount}, nil
}

func (c *casinoServer) GetTokenBalance(_ context.Context, user *commonpb.User) (*casinopb.Tokens, error) {
	log.Printf("GetTokenBalance invoked with user %v\n", user)

	usrID := userID(user.GetId())
	return &casinopb.Tokens{Count: c.userToTokens[usrID]}, nil
}

func (c *casinoServer) GetPayments(user *commonpb.User, stream casinopb.Casino_GetPaymentsServer) error {
	log.Printf("GetPayments invoked with user %v", user)

	usrID := userID(user.GetId())
	payments := c.userToPayments[usrID]
	for _, payment := range payments {
		err := stream.Send(&commonpb.Payment{
			User:   user,
			Amount: payment,
		})
		if err != nil {
			return fmt.Errorf("failed sending payment through stream: %w", err)
		}
	}

	return nil
}

func (c *casinoServer) GetPaymentStatement(ctx context.Context, user *commonpb.User) (*commonpb.PaymentStatement, error) {
	log.Printf("GetPaymentStatement invoked with user %v\n", user)

	stream, err := paymentStatementsClient.CreateStatement(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't create payment statements stream: %w", err)
	}

	usrID := userID(user.GetId())
	payments := c.userToPayments[usrID]
	for _, payment := range payments {
		err := stream.Send(&commonpb.Payment{User: user, Amount: payment})
		if err != nil {
			return nil, fmt.Errorf("failed sending payment to payments_statements: %w", err)
		}
	}

	return stream.CloseAndRecv()
}

func (c *casinoServer) Gamble(stream casinopb.Casino_GambleServer) error {
	log.Println("Gamble invoked...")

	errc := make(chan error, 2)
	go iterateStreamWithHandler(errc, stream, c.handleUserGamblingAction)
	go iterateStreamWithHandler(errc, stream, c.incrementAndSendStockPrice)

	err := <-errc
	log.Println("Gambling ending with err " + err.Error())

	return err
}

func iterateStreamWithHandler(errc chan error, stream casinopb.Casino_GambleServer, handler streamHandler) {
	for {
		select {
		case <-errc:
			return
		default:
		}

		err := handler(stream)
		if err != nil {
			errc <- err
			break
		}
	}
}

func (c *casinoServer) handleUserGamblingAction(stream casinopb.Casino_GambleServer) error {
	panic("not implemented")
}

func (c *casinoServer) incrementAndSendStockPrice(stream casinopb.Casino_GambleServer) error {
	panic("not implemented")
}
