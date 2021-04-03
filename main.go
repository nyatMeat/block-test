package main

import (
	"fmt"

	blockchain "github.com/nyatmeat/block-test/blockchain"
)

func main(){

	bc:= blockchain.NewBlockChain()
	bc.AddBlock("Send 1 block to Ivan")
	bc.AddBlock("Send 2 block to Ivan")

	for _, block := range bc.blocks{
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}