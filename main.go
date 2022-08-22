package main

import (
	"fmt"
	"github.com/Vvaradinov/evmos-erc20/deployer"
	"github.com/Vvaradinov/evmos-erc20/erc20_client"
	"os"
)

// Main entry point for the commands available for this program
// It parses the command line arguments and calls the appropriate function.
// Available commands: deployContract, balanceOf, transfer
func main() {
	argsWithoutProg := os.Args[1:]

	// Deploy contract command
	if argsWithoutProg[0] == "deployContract" {
		_, err := deployer.DeployContract(argsWithoutProg[1])
		if err != nil {
			fmt.Println("Error deploying contract: ", err)
		}
	}
	// Transfer tokens command
	if argsWithoutProg[0] == "transfer" {
		_, err := erc20_client.Transfer(argsWithoutProg[1], argsWithoutProg[2], argsWithoutProg[3], argsWithoutProg[4])
		if err != nil {
			fmt.Println("Error transferring tokens: ", err)
		}
	}
	// Query token balance command
	if argsWithoutProg[0] == "balanceOf" {
		_, err := erc20_client.QueryBalance(argsWithoutProg[1], argsWithoutProg[2])
		if err != nil {
			fmt.Println("Error querying balance: ", err)
		}
	}

}
