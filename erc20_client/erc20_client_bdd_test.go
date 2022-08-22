package erc20_client

import (
	vladcoin "github.com/Vvaradinov/evmos-erc20/contracts"
	"github.com/Vvaradinov/evmos-erc20/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/big"
	"testing"
	"time"
)

// Some constant values
// Replace with the values from running the utility script keys.utils.sh
const vladWalletPK = "9C5D6AB6AF650D0DB0DEB2347E50B72E0B4AF1E83D61589253BCE92C69192072"
const nickWalletPK = "FC3C7A7A4BA2B789C8149D949A8F34FEE9C0BBE4E3C42BAFB470D1323F397D80"
const vladWalletAddr = "9D8D9B796C20CEBF62337FF3A4019DF0917FCD41"
const nickWalletAddr = "93C54D4F2DD49FF712EE3C7CABF1634926CA046F"

// Some variables to be initialized before the tests
var tokenTransferAmount = big.NewInt(50000000)
var tooLargeTokenTransferAmount = big.NewInt(5000000000000000000)
var auth *bind.TransactOpts
var client *ethclient.Client
var contractAddr common.Address

func TestERC20(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ERC20 Suite")
}

// BeforeEach is a Ginkgo hook that is run before each test in a suite.
// Here we initialize all the required variables and objects needed for the tests.
// Obtaining client and auth, deploying contract and waiting for the block to be confirmed.
var _ = BeforeSuite(func() {
	auth, client, _ = utils.ObtainClientAndTxSigner("http://localhost:8545", vladWalletPK)
	contractAddr, _, _, _ = vladcoin.DeployVladToken(auth, client)
	time.Sleep(time.Second * 5)
})

// Close client after the tests are done
var _ = AfterSuite(func() {
	client.Close()
})

// BDD tests for query balance functionality of ERC20
var _ = Describe("Query Balance ERC20", func() {
	Context("Query the balance of Vlad and Nick after initialization", func() {
		It("Should return a balance of 10000000000000000000000000000 for Vlad", func() {
			vladBalance, _ := QueryBalance(contractAddr.Hex(), vladWalletAddr)
			Expect(vladBalance.String()).To(Equal("10000000000000000000000000000"))
		})
		It("Should return a balance of 0 for Nick", func() {
			nickBalance, _ := QueryBalance(contractAddr.Hex(), nickWalletAddr)
			Expect(nickBalance.String()).To(Equal("0"))
		})

	})
})

// BDD tests for transfer functionality of ERC20 tokens
var _ = Describe("Transfer ERC20", func() {
	// Case when the sender user has enough balance to transfer
	Context("User has sufficient balance to transfer", func() {
		BeforeEach(func() {
			// Check if the transfer transaction is successful
			status, _ := Transfer(contractAddr.Hex(), vladWalletPK, nickWalletAddr, tokenTransferAmount.String())
			Expect(status).To(Not(Equal(nil)))
		})
		It("Should increase the balance of the receiver and decrease the balance of the sender", func() {
			nickBalance, _ := QueryBalance(contractAddr.Hex(), nickWalletAddr)
			vladBalance, _ := QueryBalance(contractAddr.Hex(), vladWalletAddr)

			Expect(nickBalance.String()).To(Equal("50000000"))
			Expect(vladBalance.String()).To(Equal("9999999999999999999950000000"))
		})
	})
	// Case when the sender user does not have enough balance to transfer
	Context("User does not have sufficient balance to transfer", func() {
		BeforeEach(func() {
			// Check if the transfer transaction is successful
			status, _ := Transfer(contractAddr.Hex(), nickWalletPK, vladWalletAddr, tooLargeTokenTransferAmount.String())
			Expect(status).To(Not(Equal(nil)))
		})
		It("Should not decrease the balance of the sender and not increase the balance of the receiver", func() {
			vladBalance, _ := QueryBalance(contractAddr.Hex(), vladWalletAddr)
			nickBalance, _ := QueryBalance(contractAddr.Hex(), nickWalletAddr)

			Expect(vladBalance.String()).To(Equal("9999999999999999999950000000"))
			Expect(nickBalance.String()).To(Equal("50000000"))
		})
	})
})
