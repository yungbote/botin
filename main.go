package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	if err != nil {
		log.Fatal(err)
	}
	block, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current block: %d\n", block)
}
