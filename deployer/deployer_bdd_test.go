package deployer

import (
	"context"
	vladcoin "github.com/Vvaradinov/evmos-erc20/contracts"
	"github.com/Vvaradinov/evmos-erc20/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

// Replace with the values from running the utility script keys.utils.sh
const vladWalletPK = "9C5D6AB6AF650D0DB0DEB2347E50B72E0B4AF1E83D61589253BCE92C69192072"

var auth *bind.TransactOpts
var client *ethclient.Client

func TestERC20Deployer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ERC20 Deployer Suite")
}

// Obtaining client and auth, deploying contract and waiting for the block to be confirmed.
var _ = BeforeSuite(func() {
	auth, client, _ = utils.ObtainClientAndTxSigner("http://localhost:8545", vladWalletPK)
})

// Close client after the tests are done
var _ = AfterSuite(func() {
	client.Close()
})

// BDD tests for deploying the ERC20 contract and making sure it exists
var _ = Describe("Deploy ERC20", func() {
	Describe("Deploy ERC20 Smart Contract", func() {
		Context("Deploy to local EVM node", func() {
			// Case where we deploy the contract and wait to confirm it has been added to the next block
			It("Waits for the contract to be deployed and returns the Bytecode", func() {
				// Deploy contract
				addr, _, _, _ := vladcoin.DeployVladToken(auth, client)
				// Wait for tx to be verified and mined
				time.Sleep(time.Second * 3)
				// Check if the bytecode for the contract exists
				bytecode, _ := client.CodeAt(context.Background(), addr, nil)
				Expect(len(bytecode)).Should(BeNumerically(">", 0))
			})
			// Case where the contract is deployed, but we don't wait for confirmation
			It("Doesn't wait for the contract to be deployed and the Bytecode is empty", func() {
				// Deploy contract
				addr, _, _, _ := vladcoin.DeployVladToken(auth, client)
				// Check if the bytecode for the contract exists
				bytecode, _ := client.CodeAt(context.Background(), addr, nil)
				Expect(len(bytecode)).Should(BeNumerically("==", 0))
			})
		})
	})
})
