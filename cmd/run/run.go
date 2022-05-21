package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/kokukuma/money_box/contract"
)

const (
	// rpcProviderURL = "https://polygon-rpc.com"
	rpcProviderURL = "https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11"

	contractAddress = "8b6F44eB598e240d8608b42cc04EBdBfA360BB09"
)

func main() {
	ctx := context.Background()

	client, err := ethclient.Dial(rpcProviderURL)
	if err != nil {
		fmt.Println("failed to get client", err)
		return
	}

	auth, err := GetAuth()
	if err != nil {
		fmt.Println("failed to get auth", err)
		return
	}
	defer client.Close()

	if err := MoneyBox(ctx, client, auth); err != nil {
		fmt.Println(err)
	}
	//
	// if err := CoinMint(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }
	//
	// if err := CoinSend(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }
	//
	// if err := CoinGet(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }
}

func MoneyBox(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := contract.NewSample(common.HexToAddress(contractAddress), client)
	if err != nil {
		return err
	}

	gasPrice, err := GasPrice(ctx, client)
	if err != nil {
		return err
	}
	auth.GasPrice = gasPrice

	// opts *bind.TransactOpts, x *big.Int
	tx, err := instance.Set(auth, big.NewInt(1))
	if err != nil {
		return err
	}
	fmt.Printf("Set deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	val, err := instance.Get(&bind.CallOpts{
		Pending: false,
		From:    auth.From,
	})
	if err != nil {
		return err
	}
	fmt.Println(val)

	return nil
}

func GetAuth() (*bind.TransactOpts, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	keystore, err := os.Open(os.Getenv("KEYSTORE"))
	if err != nil {
		return nil, err
	}
	defer keystore.Close()

	keystorepass := os.Getenv("KEYSTOREPASS")
	return bind.NewTransactor(keystore, keystorepass)
}

func GasPrice(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	address := common.HexToAddress(accountAddress)
	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, err

	}
	fmt.Printf("Balance: %d\n", balance)

	return client.SuggestGasPrice(context.Background())
}
