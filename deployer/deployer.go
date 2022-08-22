package deployer

import (
	"context"
	"fmt"
	vladcoin "github.com/Vvaradinov/evmos-erc20/contracts"
	"github.com/Vvaradinov/evmos-erc20/utils"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"time"
)

func DeployContract(senderPK string) (*common.Address, error) {
	// Obtain a client and a transaction signer from utils
	auth, client, err := utils.ObtainClientAndTxSigner("http://localhost:8545", senderPK)
	if err != nil {
		return nil, err
	}

	// Deploy the token contract using our Go bindings
	addr, tx, _, err := vladcoin.DeployVladToken(auth, client)
	if err != nil {
		return nil, err
	}

	// Display the contract address and the transaction hash
	fmt.Println("The ERC20 contract address: ", addr.Hex())
	fmt.Println("Transaction hash: ", tx.Hash().Hex())

	// Sleep for 5 seconds to ensure the transaction has been completed and the block has been mined
	time.Sleep(time.Second * 5)

	// Check the transaction receipt to ensure the contract has been deployed
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatalf("Could not get the receipt for the transaction: %v", err)
		return nil, err
	}

	// Check the status of the deploy transaction
	status := receipt.Status
	if status != 1 {
		log.Fatalf("Deploy deployment failed for contract: %s\n", addr.Hex())
	}

	return &addr, nil
}
