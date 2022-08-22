package erc20_client

import (
	"context"
	"fmt"
	vladcoin "github.com/Vvaradinov/evmos-erc20/contracts"
	"github.com/Vvaradinov/evmos-erc20/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"time"
)

// QueryBalance A simple query balance function
// It queries the balance of the given address and prints it to the console.
func QueryBalance(contractAddr string, addr string) (*big.Int, error) {
	// Check for empty addresses
	if contractAddr == "" || addr == "" {
		return nil, fmt.Errorf("contract address and address must be provided")
	}

	// Obtain client
	client, _ := utils.ObtainClient("http://localhost:8545")

	// Convert address from string to Address
	tokenAddress := common.HexToAddress(contractAddr)
	originAddress := common.HexToAddress(addr)

	// Initialize the token contract from our generated Go contract bindings
	contract, err := vladcoin.NewVladToken(tokenAddress, client)
	if err != nil {
		return nil, err
	}

	// Query the balance of the given address
	balance, err := contract.BalanceOf(nil, originAddress)
	if err != nil {
		return nil, err
	}

	// Print the balance to the console
	fmt.Printf("Balance: %s\n", balance.String())
	return balance, nil
}

// Transfer A simple transfer function
// It transfers the given amount of tokens from the given address to the given address.
// It also prints the transaction hash to the console.
func Transfer(contractAddr string, fromPK string, toAddr string, amount string) (*uint64, error) {
	if contractAddr == "" || fromPK == "" || toAddr == "" {
		return nil, fmt.Errorf("contract address, from address and to address must be provided")
	}

	// Get the client and transaction signer
	auth, client, err := utils.ObtainClientAndTxSigner("http://localhost:8545", fromPK)
	if err != nil {
		return nil, err
	}

	// Convert private key to ECDSA format
	ecdsaPrivateKey, err := crypto.HexToECDSA(fromPK)
	if err != nil {
		return nil, err
	}

	// Convert the addresses from string to Address
	fromAddress := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)
	toAddress := common.HexToAddress(toAddr)
	tokenAddress := common.HexToAddress(contractAddr)

	// Converts the amount to integer format
	// Returns a custom error if it fails
	// TODO: This can be tested with a unit test
	amountInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, fmt.Errorf("could not convert amount to integer")
	}

	// Build the transaction data from the transfer amount and receiver address
	txData := utils.BuildTransferTxData(*amountInt, toAddress)

	// Build a call message for the transfer function to be executed on the EVM
	callMsg := ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: txData,
	}

	// Estimate the gas of the transaction as closely as possible.
	// Replace the gas limit in the transaction signer.
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return nil, err
	}
	auth.GasLimit = gasLimit

	// Initialize the token contract from our generated Go contract bindings
	tokenContract, err := vladcoin.NewVladToken(tokenAddress, client)
	if err != nil {
		return nil, err
	}

	// Transfer the tokens
	tx, err := tokenContract.Transfer(auth, toAddress, amountInt)
	if err != nil {
		return nil, err
	}

	// Sleep for 2 seconds to ensure the transfer is recorded, so we can check the receipt and the new balances
	time.Sleep(time.Second * 2)

	// Get the balances after the transaction
	if err != nil {
		return nil, err
	}

	// Check the receipt of the transaction and it's status
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, err
	}

	// Print the status of the transaction and the transaction hash to the console
	fmt.Printf("Transaction status: %v \n", receipt.Status)
	fmt.Printf("Transaction hash: %s\n", tx.Hash().String())

	return &receipt.Status, nil

}
