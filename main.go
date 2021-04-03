package main

import (
	blockchain "github.com/nyatmeat/block-test/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()
	defer bc.CloseDB()

	cli := blockchain.NewCLI(bc)
	cli.Run()
}
