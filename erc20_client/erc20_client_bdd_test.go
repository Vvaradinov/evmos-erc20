package erc20_client

import (
	"fmt"
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
const vladWalletPK = "37B578A789DF7A4D842A5F2934363FF811EDAFBC8F855880C2F3D9B20F757EF7"
const nickWalletPK = "61B3FC80506EF07D526917D3E3495E60946FDF3C79F3FEBD8CE125B150B2F796"
const vladWalletAddr = "9920E564CB9CDC444218B77B01A01CC6F92A9B2B"
const nickWalletAddr = "384A3B4885A891246F19315C7DFB00A3AC3ADB74"

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
	fmt.Println("The new contract address is: ", contractAddr.Hex())
	time.Sleep(time.Second * 5)
})

// Close client after the tests are done
var _ = AfterSuite(func() {
	client.Close()
})

var _ = Describe("ERC20 Token Operations", func() {
	// Case when the contract is first initialized and the balances for Nick are 0 and Vlad holds the entire token supply
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
