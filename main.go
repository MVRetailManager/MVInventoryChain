package main

import (
	"time"

	blockchainPKG "github.com/MVRetailManager/MVInventoryChain/blockchain"
	logging "github.com/MVRetailManager/MVInventoryChain/logging"
)

var (
	bc blockchainPKG.Blockchain
)

func init() {
	logging.SetupLogger()

	logging.InfoLogger.Println("Starting MVInventoryChain...")
}

func main() {
	genesisBlock := initGenesis()
	genesisBlock.Mine()
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{
			{
				Inputs: []blockchainPKG.Output{bc.Blocks[0].Transaction[0].Outputs[1]},
				Outputs: []blockchainPKG.Output{
					{
						Index:   0,
						Address: "Alice",
						Value:   2,
					},
				},
			},
		},
	)

	block.Mine()
	if bc.AddBlock(*block) != nil {
		logging.ErrorLogger.Println("Error adding block")
	}
}

func initGenesis() blockchainPKG.Block {
	return blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction: []blockchainPKG.Transaction{
			{
				Inputs: make([]blockchainPKG.Output, 0),
				Outputs: []blockchainPKG.Output{
					{
						Index:   0,
						Address: "Alice",
						Value:   30,
					},
					{
						Index:   1,
						Address: "Bob",
						Value:   7,
					},
				},
			},
		},
	}
}
