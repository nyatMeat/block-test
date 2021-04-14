package main

import (
	blockchain "github.com/nyatmeat/block-test/blockchain"
)

func main() {
	cli := blockchain.CLI{}
	cli.Run()
}
