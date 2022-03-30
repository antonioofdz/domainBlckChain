package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"code.ual/backend/internal/contracts"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x411A42CbeaE8F101508eAB8afE2D5c2d7C6e564f")
	caller, err := contracts.NewContractCaller(address, client)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := caller.HelloWorld(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded", resp)
}
