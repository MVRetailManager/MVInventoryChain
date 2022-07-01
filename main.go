package main

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	bc            Blockchain
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	warningLogger *log.Logger
	blocksLogger  *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		errorLogger.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	infoLogger = log.New(multiWriter, "INFO:		", log.LstdFlags|log.Lshortfile|log.LUTC|log.Lmicroseconds)
	warningLogger = log.New(multiWriter, "WARNING:	", log.LstdFlags|log.Lshortfile|log.LUTC|log.Lmicroseconds)
	errorLogger = log.New(multiWriter, "ERROR:		", log.LstdFlags|log.Lshortfile|log.LUTC|log.Lmicroseconds)
	blocksLogger = log.New(multiWriter, "BLOCKS:		", log.LstdFlags|log.Lshortfile|log.LUTC|log.Lmicroseconds)

	infoLogger.Println("Starting MVInventoryChain")
}

func main() {
	genesisBlock := initGenesis()
	genesisBlock.mine()
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{
			{
				inputs: []Output{bc.blocks[0].transaction[0].outputs[1]},
				outputs: []Output{
					{
						index:   0,
						address: "Alice",
						value:   2,
					},
				},
			},
		},
	)

	block.mine()
	if bc.addBlock(*block) != nil {
		errorLogger.Println("Error adding block")
	}
}

func initGenesis() Block {
	return Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction: []Transaction{
			{
				inputs: make([]Output, 0),
				outputs: []Output{
					{
						index:   0,
						address: "Alice",
						value:   30,
					},
					{
						index:   1,
						address: "Bob",
						value:   7,
					},
				},
			},
		},
	}
}
