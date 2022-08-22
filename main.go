package main

import (
	"fmt"
	"github.com/Vvaradinov/evmos-erc20/deployer"
	"github.com/Vvaradinov/evmos-erc20/erc20_client"
	"os"
)

func main() {
	const contractAddr = "0x0f2832522dC01C8dF87c6aE007EC10d06a7335db"
	const userVladPK = "D55B4BD2106C8691298B8A110F734CE7353A2CFBCECFD698028DEF752451AF6E"
	const userVlad = "0x161e373958716Dc16Efb4659CE80f926EC0399a2"
	const userNick = "0x58BA2585E068720D989FC2DC2DF991C5DA3DA7C0"
	const erc20Contract = "0x4022411E2bd4b19E7EFc474Bb620db2209C7F5Fa"

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
