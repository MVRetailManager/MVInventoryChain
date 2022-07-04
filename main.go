package main

import (
	"os"

	"github.com/MVRetailManager/MVInventoryChain/cli"
	logging "github.com/MVRetailManager/MVInventoryChain/logging"
)

func init() {
	logging.SetupLogger()

	logging.InfoLogger.Println("Starting MVInventoryChain...")
}

func main() {
	defer os.Exit(0)
	cmdline := cli.CLI{}
	cmdline.Run()
	/*
		block := blockchainPKG.NewBlock(
			1,
			time.Now().UTC().UnixNano(),
			genesisBlock.Hash,
			1,
			[]blockchainPKG.Transaction{
				{
					Inputs: []blockchainPKG.Output{
						{
							Index:   0,
							Address: "Larry",
							Value:   200,
						},
					},
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
		}*/
}
