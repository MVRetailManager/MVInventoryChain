package main

import (
	"fmt"
	"time"
)

func main() {
	var bc Blockchain

	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{
			{
				inputs: []Output{
					{
						index:   0,
						address: "Bob",
						value:   100000,
					},
				},
				outputs: []Output{
					{
						index:   0,
						address: "Alice",
						value:   1,
					},
				},
			},
		},
	)

	block.mine()

	//fmt.Println(block)

	bc.addBlock(*block)

	fmt.Println(bc)
}
