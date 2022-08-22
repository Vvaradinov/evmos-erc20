// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VladToken

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VladTokenMetaData contains all meta data concerning the VladToken contract.
var VladTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600881526020017f566c6164636f696e0000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f564c41440000000000000000000000000000000000000000000000000000000081525081600390816200008f9190620004c4565b508060049081620000a19190620004c4565b505050620000c2336b204fce5e3e25026110000000620000c860201b60201c565b620006c6565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036200013a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000131906200060c565b60405180910390fd5b6200014e600083836200024060201b60201c565b80600260008282546200016291906200065d565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620001b991906200065d565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620002209190620006a9565b60405180910390a36200023c600083836200024560201b60201c565b5050565b505050565b505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002cc57607f821691505b602082108103620002e257620002e162000284565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200034c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200030d565b6200035886836200030d565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003a56200039f620003998462000370565b6200037a565b62000370565b9050919050565b6000819050919050565b620003c18362000384565b620003d9620003d082620003ac565b8484546200031a565b825550505050565b600090565b620003f0620003e1565b620003fd818484620003b6565b505050565b5b81811015620004255762000419600082620003e6565b60018101905062000403565b5050565b601f82111562000474576200043e81620002e8565b6200044984620002fd565b8101602085101562000459578190505b620004716200046885620002fd565b83018262000402565b50505b505050565b600082821c905092915050565b6000620004996000198460080262000479565b1980831691505092915050565b6000620004b4838362000486565b9150826002028217905092915050565b620004cf826200024a565b67ffffffffffffffff811115620004eb57620004ea62000255565b5b620004f78254620002b3565b6200050482828562000429565b600060209050601f8311600181146200053c576000841562000527578287015190505b620005338582620004a6565b865550620005a3565b601f1984166200054c86620002e8565b60005b8281101562000576578489015182556001820191506020850194506020810190506200054f565b8683101562000596578489015162000592601f89168262000486565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000620005f4601f83620005ab565b91506200060182620005bc565b602082019050919050565b600060208201905081810360008301526200062781620005e5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200066a8262000370565b9150620006778362000370565b92508282019050808211156200069257620006916200062e565b5b92915050565b620006a38162000370565b82525050565b6000602082019050620006c0600083018462000698565b92915050565b61123880620006d66000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610b15565b60405180910390f35b6100e660048036038101906100e19190610bd0565b610308565b6040516100f39190610c2b565b60405180910390f35b61010461032b565b6040516101119190610c55565b60405180910390f35b610134600480360381019061012f9190610c70565b610335565b6040516101419190610c2b565b60405180910390f35b610152610364565b60405161015f9190610cdf565b60405180910390f35b610182600480360381019061017d9190610bd0565b61036d565b60405161018f9190610c2b565b60405180910390f35b6101b260048036038101906101ad9190610cfa565b6103a4565b6040516101bf9190610c55565b60405180910390f35b6101d06103ec565b6040516101dd9190610b15565b60405180910390f35b61020060048036038101906101fb9190610bd0565b61047e565b60405161020d9190610c2b565b60405180910390f35b610230600480360381019061022b9190610bd0565b6104f5565b60405161023d9190610c2b565b60405180910390f35b610260600480360381019061025b9190610d27565b610518565b60405161026d9190610c55565b60405180910390f35b60606003805461028590610d96565b80601f01602080910402602001604051908101604052809291908181526020018280546102b190610d96565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b60008061031361059f565b90506103208185856105a7565b600191505092915050565b6000600254905090565b60008061034061059f565b905061034d858285610770565b6103588585856107fc565b60019150509392505050565b60006012905090565b60008061037861059f565b905061039981858561038a8589610518565b6103949190610df6565b6105a7565b600191505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546103fb90610d96565b80601f016020809104026020016040519081016040528092919081815260200182805461042790610d96565b80156104745780601f1061044957610100808354040283529160200191610474565b820191906000526020600020905b81548152906001019060200180831161045757829003601f168201915b5050505050905090565b60008061048961059f565b905060006104978286610518565b9050838110156104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d390610e9c565b60405180910390fd5b6104e982868684036105a7565b60019250505092915050565b60008061050061059f565b905061050d8185856107fc565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610616576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060d90610f2e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610685576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067c90610fc0565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107639190610c55565b60405180910390a3505050565b600061077c8484610518565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107f657818110156107e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107df9061102c565b60405180910390fd5b6107f584848484036105a7565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361086b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610862906110be565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d190611150565b60405180910390fd5b6108e5838383610a7b565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561096b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610962906111e2565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109fe9190610df6565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a629190610c55565b60405180910390a3610a75848484610a80565b50505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610abf578082015181840152602081019050610aa4565b60008484015250505050565b6000601f19601f8301169050919050565b6000610ae782610a85565b610af18185610a90565b9350610b01818560208601610aa1565b610b0a81610acb565b840191505092915050565b60006020820190508181036000830152610b2f8184610adc565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b6782610b3c565b9050919050565b610b7781610b5c565b8114610b8257600080fd5b50565b600081359050610b9481610b6e565b92915050565b6000819050919050565b610bad81610b9a565b8114610bb857600080fd5b50565b600081359050610bca81610ba4565b92915050565b60008060408385031215610be757610be6610b37565b5b6000610bf585828601610b85565b9250506020610c0685828601610bbb565b9150509250929050565b60008115159050919050565b610c2581610c10565b82525050565b6000602082019050610c406000830184610c1c565b92915050565b610c4f81610b9a565b82525050565b6000602082019050610c6a6000830184610c46565b92915050565b600080600060608486031215610c8957610c88610b37565b5b6000610c9786828701610b85565b9350506020610ca886828701610b85565b9250506040610cb986828701610bbb565b9150509250925092565b600060ff82169050919050565b610cd981610cc3565b82525050565b6000602082019050610cf46000830184610cd0565b92915050565b600060208284031215610d1057610d0f610b37565b5b6000610d1e84828501610b85565b91505092915050565b60008060408385031215610d3e57610d3d610b37565b5b6000610d4c85828601610b85565b9250506020610d5d85828601610b85565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610dae57607f821691505b602082108103610dc157610dc0610d67565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e0182610b9a565b9150610e0c83610b9a565b9250828201905080821115610e2457610e23610dc7565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000610e86602583610a90565b9150610e9182610e2a565b604082019050919050565b60006020820190508181036000830152610eb581610e79565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000610f18602483610a90565b9150610f2382610ebc565b604082019050919050565b60006020820190508181036000830152610f4781610f0b565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000610faa602283610a90565b9150610fb582610f4e565b604082019050919050565b60006020820190508181036000830152610fd981610f9d565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611016601d83610a90565b915061102182610fe0565b602082019050919050565b6000602082019050818103600083015261104581611009565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006110a8602583610a90565b91506110b38261104c565b604082019050919050565b600060208201905081810360008301526110d78161109b565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b600061113a602383610a90565b9150611145826110de565b604082019050919050565b600060208201905081810360008301526111698161112d565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006111cc602683610a90565b91506111d782611170565b604082019050919050565b600060208201905081810360008301526111fb816111bf565b905091905056fea26469706673582212200c78355173e04cd98e77f53de90d40378e544f006c2b24af39494da0d241334b64736f6c63430008100033",
}

// VladTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use VladTokenMetaData.ABI instead.
var VladTokenABI = VladTokenMetaData.ABI

// VladTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VladTokenMetaData.Bin instead.
var VladTokenBin = VladTokenMetaData.Bin

// DeployVladToken deploys a new Ethereum contract, binding an instance of VladToken to it.
func DeployVladToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VladToken, error) {
	parsed, err := VladTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VladTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VladToken{VladTokenCaller: VladTokenCaller{contract: contract}, VladTokenTransactor: VladTokenTransactor{contract: contract}, VladTokenFilterer: VladTokenFilterer{contract: contract}}, nil
}

// VladToken is an auto generated Go binding around an Ethereum contract.
type VladToken struct {
	VladTokenCaller     // Read-only binding to the contract
	VladTokenTransactor // Write-only binding to the contract
	VladTokenFilterer   // Log filterer for contract events
}

// VladTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type VladTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VladTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VladTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VladTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VladTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VladTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VladTokenSession struct {
	Contract     *VladToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VladTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VladTokenCallerSession struct {
	Contract *VladTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VladTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VladTokenTransactorSession struct {
	Contract     *VladTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VladTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type VladTokenRaw struct {
	Contract *VladToken // Generic contract binding to access the raw methods on
}

// VladTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VladTokenCallerRaw struct {
	Contract *VladTokenCaller // Generic read-only contract binding to access the raw methods on
}

// VladTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VladTokenTransactorRaw struct {
	Contract *VladTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVladToken creates a new instance of VladToken, bound to a specific deployed contract.
func NewVladToken(address common.Address, backend bind.ContractBackend) (*VladToken, error) {
	contract, err := bindVladToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VladToken{VladTokenCaller: VladTokenCaller{contract: contract}, VladTokenTransactor: VladTokenTransactor{contract: contract}, VladTokenFilterer: VladTokenFilterer{contract: contract}}, nil
}

// NewVladTokenCaller creates a new read-only instance of VladToken, bound to a specific deployed contract.
func NewVladTokenCaller(address common.Address, caller bind.ContractCaller) (*VladTokenCaller, error) {
	contract, err := bindVladToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VladTokenCaller{contract: contract}, nil
}

// NewVladTokenTransactor creates a new write-only instance of VladToken, bound to a specific deployed contract.
func NewVladTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*VladTokenTransactor, error) {
	contract, err := bindVladToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VladTokenTransactor{contract: contract}, nil
}

// NewVladTokenFilterer creates a new log filterer instance of VladToken, bound to a specific deployed contract.
func NewVladTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*VladTokenFilterer, error) {
	contract, err := bindVladToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VladTokenFilterer{contract: contract}, nil
}

// bindVladToken binds a generic wrapper to an already deployed contract.
func bindVladToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VladTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VladToken *VladTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VladToken.Contract.VladTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VladToken *VladTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VladToken.Contract.VladTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VladToken *VladTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VladToken.Contract.VladTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VladToken *VladTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VladToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VladToken *VladTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VladToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VladToken *VladTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VladToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VladToken *VladTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VladToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VladToken *VladTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VladToken.Contract.Allowance(&_VladToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VladToken *VladTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VladToken.Contract.Allowance(&_VladToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VladToken *VladTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VladToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VladToken *VladTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VladToken.Contract.BalanceOf(&_VladToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VladToken *VladTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VladToken.Contract.BalanceOf(&_VladToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VladToken *VladTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VladToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VladToken *VladTokenSession) Decimals() (uint8, error) {
	return _VladToken.Contract.Decimals(&_VladToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VladToken *VladTokenCallerSession) Decimals() (uint8, error) {
	return _VladToken.Contract.Decimals(&_VladToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VladToken *VladTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VladToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VladToken *VladTokenSession) Name() (string, error) {
	return _VladToken.Contract.Name(&_VladToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VladToken *VladTokenCallerSession) Name() (string, error) {
	return _VladToken.Contract.Name(&_VladToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VladToken *VladTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VladToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VladToken *VladTokenSession) Symbol() (string, error) {
	return _VladToken.Contract.Symbol(&_VladToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VladToken *VladTokenCallerSession) Symbol() (string, error) {
	return _VladToken.Contract.Symbol(&_VladToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VladToken *VladTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VladToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VladToken *VladTokenSession) TotalSupply() (*big.Int, error) {
	return _VladToken.Contract.TotalSupply(&_VladToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VladToken *VladTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _VladToken.Contract.TotalSupply(&_VladToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VladToken *VladTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VladToken *VladTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.Approve(&_VladToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VladToken *VladTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.Approve(&_VladToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VladToken *VladTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VladToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VladToken *VladTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.DecreaseAllowance(&_VladToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VladToken *VladTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.DecreaseAllowance(&_VladToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VladToken *VladTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VladToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VladToken *VladTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.IncreaseAllowance(&_VladToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VladToken *VladTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.IncreaseAllowance(&_VladToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VladToken *VladTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VladToken *VladTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.Transfer(&_VladToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VladToken *VladTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.Transfer(&_VladToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VladToken *VladTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VladToken *VladTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.TransferFrom(&_VladToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VladToken *VladTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VladToken.Contract.TransferFrom(&_VladToken.TransactOpts, from, to, amount)
}

// VladTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VladToken contract.
type VladTokenApprovalIterator struct {
	Event *VladTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VladTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VladTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VladTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VladTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VladTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VladTokenApproval represents a Approval event raised by the VladToken contract.
type VladTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VladToken *VladTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VladTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VladToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VladTokenApprovalIterator{contract: _VladToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VladToken *VladTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VladTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VladToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VladTokenApproval)
				if err := _VladToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VladToken *VladTokenFilterer) ParseApproval(log types.Log) (*VladTokenApproval, error) {
	event := new(VladTokenApproval)
	if err := _VladToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VladTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VladToken contract.
type VladTokenTransferIterator struct {
	Event *VladTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VladTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VladTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VladTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VladTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VladTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VladTokenTransfer represents a Transfer event raised by the VladToken contract.
type VladTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VladToken *VladTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VladTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VladToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VladTokenTransferIterator{contract: _VladToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VladToken *VladTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VladTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VladToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VladTokenTransfer)
				if err := _VladToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VladToken *VladTokenFilterer) ParseTransfer(log types.Log) (*VladTokenTransfer, error) {
	event := new(VladTokenTransfer)
	if err := _VladToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
