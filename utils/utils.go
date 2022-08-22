package utils

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

// ObtainClient Initializes a Client instance and connects to a node
func ObtainClient(rpcUrl string) (*ethclient.Client, error) {
	// Connect to a node with the RPC URL
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// ObtainClientAndTxSigner Obtains the client and transaction signer
// which we will use to sign transactions and send them to the network.
func ObtainClientAndTxSigner(rpcUrl string, userPK string) (*bind.TransactOpts, *ethclient.Client, error) {
	// Connect to a node with the RPC URL
	client, err := ObtainClient(rpcUrl)
	if err != nil {
		return nil, nil, err
	}

	// Get private key for our local node test account
	privateKey, err := crypto.HexToECDSA(userPK)
	if err != nil {
		return nil, nil, err
	}

	// Get the public key from our private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// Get the public key in the form of an address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Get the suggested gas price from the client
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Get the EVM chain ID from the client
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, err
	}

	// Create a new KeyTransactor with our private key and chain ID.
	// Update necessary information for the signer
	// including nonce, gas price, and gas limit.
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(30000000000000000)
	auth.GasPrice = gasPrice
	return auth, client, nil
}

// BuildTransferTxData Builds a ERC20 Transfer transaction.
// Populates it with the receiver address and the desired amount.
// Returns a byte array.
func BuildTransferTxData(amount big.Int, toAddress common.Address) []byte {
	// Using the ERC20 transfer function signature
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	// Zero-pads byte slice to the left up to length l.
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	// Concatenates the methodID, paddedAmount, and paddedAddress
	// to the final byte array
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	return data
}
