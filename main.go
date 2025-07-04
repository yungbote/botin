package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	POOL1 = "0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"
	POOL2 = "0x397FF1542f962076d0BFE58eA045FfA2d347ACa0"
)

func main() {
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Checking prices...")
		//TODO: Check prices
		block, _ := client.BlockNumber(context.Background())
		fmt.Printf("Block: %d\n", block)
		time.Sleep(12 * time.Second)
	}
}
