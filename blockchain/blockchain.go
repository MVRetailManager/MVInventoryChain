package blockchain

import (
	"bytes"
	"time"

	bcErrors "github.com/MVRetailManager/MVInventoryChain/customErrors"
	logging "github.com/MVRetailManager/MVInventoryChain/logging"
)

type Blockchain struct {
	GenesisBlock Block
	Blocks       []Block
}

func (bc *Blockchain) NewBlockchain(genesisBlock Block) {
	bc.GenesisBlock = genesisBlock
	bc.Blocks = []Block{genesisBlock}

	logging.InfoLogger.Printf("New blockchain created with genesis block: %s", bc.GenesisBlock.Hash)
}

func (bc *Blockchain) AddBlock(block Block) error {
	if err := bc.isValidBlock(block); err != nil {
		logging.WarningLogger.Printf(err.Error())
		return err
	}

	bc.Blocks = append(bc.Blocks, block)

	return nil
}

func (bc *Blockchain) isValidBlock(block Block) error {
	if bc.isInvalidIndex(block) {
		return bcErrors.MismatchedIndex{ExpectedIndex: len(bc.Blocks), ActualIndex: block.Index}.DoError()
	}
	if bc.isAchronTimestamp(block) {
		return bcErrors.AchronologicalTimestamp{ExpectedTimestamp: bc.Blocks[len(bc.Blocks)-1].UnixTimeStamp, ActualTimestamp: block.UnixTimeStamp}.DoError()
	}
	if bc.isInvalidTimestamp(block) {
		return bcErrors.InvalidTimestamp{ExpectedTimestamp: time.Now().UTC().UnixNano(), ActualTimestamp: block.UnixTimeStamp}.DoError()
	}
	if bc.isInvalidPreviousHash(block) {
		return bcErrors.InvalidPreviousHash{ExpectedHash: bc.Blocks[len(bc.Blocks)-1].Hash, ActualHash: block.PreviousHash}.DoError()
	}
	if bc.isInvalidGenesisBlock(block) {
		return bcErrors.InvalidGenesisBlock{ExpectedFormat: make([]byte, 32), ActualFormat: bc.GenesisBlock.Hash}.DoError()
	}

	for _, transaction := range block.Transaction {
		if err := bc.isInsufficientInputValue(transaction); err != nil {
			return err
		}
		if err := bc.isInputValueLessZero(transaction); err != nil {
			return err
		}
	}

	return nil
}

func (bc *Blockchain) isInvalidIndex(block Block) bool {
	return block.Index != len(bc.Blocks)
}

func (bc *Blockchain) isAchronTimestamp(block Block) bool {
	return block.UnixTimeStamp < bc.Blocks[len(bc.Blocks)-1].UnixTimeStamp
}

func (bc *Blockchain) isInvalidTimestamp(block Block) bool {
	return block.UnixTimeStamp > time.Now().UTC().UnixNano()
}

func (bc *Blockchain) isInvalidPreviousHash(block Block) bool {
	return !bytes.Equal(block.PreviousHash, bc.Blocks[len(bc.Blocks)-1].Hash)
}

func (bc *Blockchain) isInvalidGenesisBlock(block Block) bool {
	return !bytes.Equal(bc.GenesisBlock.Hash, make([]byte, 32))
}

func (bc *Blockchain) isInsufficientInputValue(transaction Transaction) error {
	if transaction.inputValue() < transaction.outputValue() {
		return bcErrors.InsufficientInputValue{ExpectedValue: transaction.inputValue(), ActualValue: transaction.outputValue()}.DoError()
	}

	return nil
}

func (bc *Blockchain) isInputValueLessZero(transaction Transaction) error {
	for _, input := range transaction.Inputs {
		if input.Value < 0 {
			return bcErrors.InvalidInput{ExpectedValue: 0, ActualValue: input.Value}.DoError()
		}
	}

	return nil
}
