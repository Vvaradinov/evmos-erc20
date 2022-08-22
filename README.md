# Evmos ERC20 Contract
- [Summary](#summary)
- [Project Structure](#project-structure)
- [Local Node Setup](#local-node-setup)
- [ERC20 Contract](#erc20-contract-deployment)
- [Programmatic Deployment](#programatic-deployment)
- [Client Interactions](#client-interactions)
- [Running Tests](#running-tests)
- [Future Improvements](#future-improvements)

## Summary
This project aims to demonstrate how to deploy an ERC20 contract on a local node using only
[go-ethereum](https://github.com/ethereum/go-ethereum) by generating [Go bindings using abigen](https://geth.ethereum.org/docs/dapp/native-bindings#abigen-go-binding-generator).
Additionally, it provides simple client functionality that can query ERC20 balances and transfer ERC20 tokens.
Unit tests are provided to verify individual functions correctness while a
[Behaviour Driven Development (BDD)](https://medium.com/javascript-scene/behavior-driven-development-bdd-and-functional-testing-62084ad7f1f2) 
approach to tests are provided using [ginko](https://github.com/onsi/ginkgo) and [gomega](https://github.com/onsi/gomega)
to verify the behaviour of the implementation and establish an easy to maintain and extend test suite.
To demonstrate the functionality of the project, we will assume a scenario where two friends **Vlad** and **Nick** are
sending each other tokens.

## Project Structure
```
├── contracts // ERC20 contract, Golang bindings and compiled contract artifacts
│   ├── VladToken.go
│   ├── VladToken.sol
│   └── build
├── deployer // Deployer script to deploy the contract on a local node and BDD tests
│   ├── deployer.go
│   └── deployer_bdd_test.go
├── erc20-client // Erc20 client scripts to query and transfer ERC20 tokens and BDD and unit tests
│   ├── erc20_client.go
│   ├── erc20_client_bdd_test.go
│   └── erc20_client_test.go  
├── utils // Utility scripts for ethereum client connectin, tx signer generation and unit tests 
│   ├── utils.go
│   └── utils_test.go  
├── node // Files generated after running local node - genesis file, private key and node configs
│   ├── config
│   ├── data
│   └── keyring-test
├── init.sh // Script used to start a local node and create user keys     
├── keys_util.sh // Script used to easily view user Hex addresses and Private Keys     
├── main.go // Main entry point for the application where all the commands are
├── go.mod     
├── package.json
├── package-lock.json
├── README.md     
```

## Local Node Setup
### Preparing your environment
Make sure you follow the instructions in the official [Evmos documentation](https://docs.evmos.org/validators/quickstart/installation.html)
to install [Go](https://go.dev/dl/) and install the latest version of [Evmos](https://github.com/evmos/evmos/releases)
###  Node Setup
A convenience script has been prepared called `init.sh` which will initialize a local node and generate user keys.
In our case these keys are `vlad_validator` and `nick_validator`. Feel free to modify the key names and `HOME_DIR`
variable to suit your local environment. Once you have successfully started your local node make sure to keep the terminal window open.

## ERC20 Contract
### Contract file
For the ERC20 contract we will use an already established and trusted implementation from [OpenZeppelin](#https://docs.openzeppelin.com/contracts/4.x/erc20).
The token contract can be found under the `contracts` folder under the name `VladToken.sol`. 
You can install OpenZepelin contracts directly with NPM using `npm install @openzeppelin/contracts`.

The contract defines a simple ERC20 contract name **Vladcoin** with symbol **VLAD** and initial supply of `100,000,000.00` which will all be delivered to the contract initializer.

### Contract Compilation
The first step is to install the [Solidity compiler (solc)](https://docs.soliditylang.org/en/latest/installing-solidity.html).
We also need to install a tool called `abigen` for generating the ABI from a solidity smart contract.
Assuming you have Go all set up on your computer, simply run the following to install the abigen tool
```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```
After this you can compile your contract and generate Go bindings using the following commands

`solc --abi --bin VladToken.sol -o build --base-path 'path_to_node_modules'`    

This will create a `build` folder with the generated `ABI` and `bin` files. Make sure to replace 
`--base-path` with the path to your `node_modules` folder so that the compiler knows where to look for
OpenZeppelin contracts that are being imported in our ERC20 contract.

`abigen --bin=contracts/build/VladToken.bin --abi=contracts/build/VladToken.abi --out=VladToken.go`

This will compile the necessary Go file which will include the deployment methods which we will
later user to in our implementation. Now you should have a `VladToken.go` file in the `contracts` folder.


## Programmatic Deployment
In order to deploy an ERC20 contract a deployer account is needed. We can get our default `vlad_validator` account's
Private Key using a command by the Evmos daemon.

`evmosd keys unsafe-export-eth-key vlad_validator --keyring-backend test`

The response might look something like this:

`5CE50C38D30F7C95EE50ACC842ACE8D1D7C1397B991259F0AF7B0780B079F958`

This is the private key of the `vlad_validator` account which we use as the default contract deployer. Make sure to note down
your private key and keep it safe as we will need it later.

A small CLI utility has been provided to deploy the ERC20 contract using the private key. 
In a new terminal window run the following command making sure you replace the private key with the one your own.

`go run main.go deployContract deployer-private-key`

Or in our case:

`go run main.go deployContract 5CE50C38D30F7C95EE50ACC842ACE8D1D7C1397B991259F0AF7B0780B079F958`

You should now see the output in your terminal if there were no errors. Make sure you ****record** the contract address**
as we will be using it in the next section [Client Interactions](#client-interactions) to interact with the contract.

```
The ERC20 contract address:  0x22B17eAB16deEC032c41f31e8Aadfd62BE9F2863
Transaction hash:  0x33ac13ecc05a3922ebaf5d560e141a5b1e196954e9002790512f8410c50e601a
```

## Client Interactions
Now that we have successfully deployed our ERC20 smart contract we can start interacting with it.
For now the only 2 functions that we will be using are `balanceOf` and `transfer`. We will start by checking
the balance of the `vlad_validator` account which should contain the entire initial supply of **Vladcoin**

In order to transfer tokens to `nick_validator` we will need to transform the address from `bech32` format to `hex` format.
Luckily for us the Evmos daemon has a utility called `evmosd keys parse` which can be used to do this.
Before we can interact with the contract we need to get the 2 accounts we generated earlier with our `init.sh` script.
There is a utility script called `keys_util.sh` which will do the necessary conversions and give you the Hex addresses for both users
and their private keys which you can use to check the balance of the account or send a transfer transaction.

In your terminal window run the utility script 

`./keys_util.sh`

That will result in an output in the following format:

```
Vlad Hex key - 161E373958716DC16EFB4659CE80F926EC0399A2
Nick Hex key - 58BA2585E068720D989FC2DC2DF991C5DA3DA7C0
Vlad Private Key - D55B4BD2106C8691298B8A110F734CE7353A2CFBCECFD698028DEF752451AF6E
Nick Private Key - DBFA8274FCBEA07BD3EA649464E15082F4D53681D2D5C11C8297B25BF9C139A
```

Now that we have the addresses in hex format we can use the client functions to transfer some **VLAD** tokens to `nick_valdiator`.
But before that let's make sure that `vlad_validator` really has a healthy balance.

In your terminal execute the following command making sure to replace with your own address:

`go run main.go balanceOf erc20-contract-address user-hex-address`

Or in our case:

`go run main.go balanceOf 0x4022411E2bd4b19E7EFc474Bb620db2209C7F5Fa 0x161e373958716Dc16Efb4659CE80f926EC0399a2`

The result will look something like this

` Balance: 10000000000000000000000000000 `

Now that we know `vlad_validator` has some **VLAD** tokens we can transfer some to `nick_validator`.
In your terminal type the following command making sure to replace with your own address and private key:

`go run main.go transfer er20-contract-address sender-private-key receiver-hex-address amount`

Or in our case:

` go run main.go transfer 0x22B17eAB16deEC032c41f31e8Aadfd62BE9F2863 5CE50C38D30F7C95EE50ACC842ACE8D1D7C1397B991259F0AF7B0780B079F958 161E373958716DC16EFB4659CE80F926EC0399A2 50000000000`

The result will look something like this:
```
Transaction status: 1
Transaction hash: 0xeae34e82ba3b4c39359a86da694fcae8030ffb72f43e552f66c31ecfcee51ddd
```

Now we can check to see if `nick_validator` has the tokens we just transferred.
In your terminal type the following command making sure to replace with your own address:

`go run main.go balanceOf erc20-contract-address user-address`

Or in our case:

`go run main.go balanceOf 0x423DE28D5d2d223B2990cc15Dc591A9E49cC30a3 0x58BA2585E068720D989FC2DC2DF991C5DA3DA7C0`

The result will look something like this

` Balance: 50000000000 `

Just to be sure lets also check `vlad_validator`s  balance as well.

`go run main.go balanceOf 0x423DE28D5d2d223B2990cc15Dc591A9E49cC30a3 0x161E373958716DC16EFB4659CE80F926EC0399A2`

The result is:

` Balance: 9999999999999999950000000000 `


## Running Tests
There are 2 types of tests implemented in this project.
- **Table Driven Tests** that check for the correct input and the expected return values for the individual functions  `Transfer`, `QueryBalance`, `ObtainClient` and `ObtainClientAndTxSigner`
- **Behaviour Driven Devlopment (BDD)** inspired tests written with  [ginko](https://github.com/onsi/ginkgo) and [gomega](https://github.com/onsi/gomega)

**In order to run the tests you first need to replace the constant values in all BDD test files with your own uniquely generated accounts.
You can do this by executing the utility script `./keys_utils.sh`**

### Instructions
Navigate to the `deployer` module and run `go test` to begin the test suite
You should get the following output if you replaced the constant values and followed the steps on launching a local node

```
Random Seed: 1661193377

Will run 2 of 2 specs
••

Ran 2 of 2 Specs in 3.011 seconds
SUCCESS! -- 2 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
ok      github.com/Vvaradinov/evmos-erc20/deployer      3.025s
```

Navigate to the `erc20_client` module and run `go test` to begin the test suite.
You should get the following output if you replaced the constant values and followed the steps on launching a local node.

```
Random Seed: 1661193539

Will run 4 of 4 specs
Balance: 10000000000000000000000000000
•Balance: 0
•Transaction status: 1
Transaction hash: 0xc57a573ba8026c96afec92ea6c7db0ddd08174c031320bf755bdeec66abeee3f
Balance: 50000000
Balance: 9999999999999999999950000000
•Balance: 9999999999999999999950000000
Balance: 50000000
•

Ran 4 of 4 Specs in 7.027 seconds
SUCCESS! -- 4 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
ok      github.com/Vvaradinov/evmos-erc20/erc20_client  7.040s
```

## Future Improvements
- Add more test cases in the `Table Driven Tests` section
- Build a CLI utility with Cobra to interact with the contract
- Automatically inject the Private keys and Hex addresses into the BDD tests
- Subscribe to contract deploy and transfer events  order to avoid using `time.Sleep` function to ensure the transfer or deploy event is included in the block