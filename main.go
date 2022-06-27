package main

import (
	"fmt"
	"time"
)

func main() {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	blockchain := Blockchain{
		genesisBlock: genesisBlock,
		blocks:       []Block{genesisBlock},
	}

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	fmt.Println(block)

	fmt.Println(blockchain)
}

func test() string {
	return "Hello, world."
}
