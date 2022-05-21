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
	money_box "github.com/kokukuma/money_box/contract"
)

const (
	// rpcProviderURL = "https://polygon-rpc.com"
	rpcProviderURL = "https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11"

	contractAddress = "a5D4B887Ac0d118D9C2C7c75918379B438E85A00"
	accountAddress  = "7E532D0E80480eCF1b566920752840bc53556366"
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

	if err := MoneyBoxGetAmoutn(ctx, client, auth); err != nil {
		fmt.Println(err)
	}
	if err := MoneyBoxPray(ctx, client, auth); err != nil {
		fmt.Println(err)
	}
	// if err := MoneyGodCollect(ctx, client, auth); err != nil {
	// 	fmt.Println(err)
	// }
}

func MoneyBoxGetAmoutn(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := money_box.NewMoneyBox(common.HexToAddress(contractAddress), client)
	if err != nil {
		return err
	}

	val, err := instance.Balance(&bind.CallOpts{
		Pending: false,
		From:    auth.From,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Current Amount: %d\n", val)
	return nil
}

func MoneyBoxPray(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := money_box.NewMoneyBox(common.HexToAddress(contractAddress), client)
	if err != nil {
		return err
	}

	gasPrice, err := GasPrice(ctx, client)
	if err != nil {
		return err
	}
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(100)

	// opts *bind.TransactOpts, x *big.Int
	tx, err := instance.Pray(auth)
	if err != nil {
		return err
	}
	fmt.Printf("Set deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	return nil
}

func MoneyGodCollect(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	instance, err := money_box.NewMoneyBox(common.HexToAddress(contractAddress), client)
	if err != nil {
		return err
	}

	gasPrice, err := GasPrice(ctx, client)
	if err != nil {
		return err
	}
	gasPrice.Mul(gasPrice, big.NewInt(3))
	auth.GasPrice = gasPrice

	// opts *bind.TransactOpts, x *big.Int
	tx, err := instance.GodCollect(auth)
	if err != nil {
		return err
	}
	fmt.Printf("Set deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

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
