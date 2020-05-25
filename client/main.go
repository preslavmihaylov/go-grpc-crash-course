package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	casinopb "gitlab.com/preslavmihaylov/go-grpc-exercise/gen/casino"
	"google.golang.org/grpc"
)

const (
	emptyCmd            command = ""
	buyTokensCmd                = "buyTokens"
	withdrawCmd                 = "withdraw"
	tokenBalanceCmd             = "balance"
	gambleCmd                   = "gamble"
	paymentsCmd                 = "payments"
	paymentStatementCmd         = "paymentStatement"
)

var reader *bufio.Reader
var username string
var client casinopb.CasinoClient

func main() {
	var conn *grpc.ClientConn

	client, conn = setupClient()
	defer conn.Close()
	log.Println("Successfully connected to casino server...")

	username = loginUser()
	fmt.Printf("welcome back %v!\n", username)

	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.Trim(input, "\n ")
		tokens := regexp.MustCompile("[ ]+").Split(input, -1)

		cmd, args := tokens[0], tokens[1:]
		if cmd == "exit" {
			fmt.Println("Bye, bye!")
			return
		}

		res, err := processCmd(command(cmd), args)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		if res != "" {
			fmt.Println(res)
		}
	}
}

func loginUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("couldn't read username: " + err.Error())
	}

	return strings.Trim(input, "\n ")
}

func processCmd(cmd command, args []string) (string, error) {
	switch cmd {
	case emptyCmd:
		return "", nil
	case buyTokensCmd:
		if len(args) != 1 {
			return "", errors.New("invalid arguments")
		}

		tokensCnt, err := strconv.Atoi(args[0])
		if err != nil {
			return "", fmt.Errorf("couldn't parse cmd argument %s: %w", args[0], err)
		}

		return buyTokens(tokensCnt)
	case withdrawCmd:
		if len(args) != 1 {
			return "", errors.New("invalid arguments")
		}

		tokensCnt, err := strconv.Atoi(args[0])
		if err != nil {
			return "", fmt.Errorf("couldn't parse cmd argument %s: %w", args[0], err)
		}

		return withdraw(tokensCnt)
	case tokenBalanceCmd:
		if len(args) != 0 {
			return "", errors.New("invalid arguments")
		}

		return tokenBalance()
	case gambleCmd:
		return gamble()
	case paymentsCmd:
		if len(args) != 0 {
			return "", errors.New("invalid arguments")
		}

		return payments()
	case paymentStatementCmd:
		return paymentStatement()
	default:
		return "", errors.New("unknown command")
	}
}
