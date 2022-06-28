package main

import (
	"bytes"
	"fmt"
	"time"
)

type Blockchain struct {
	genesisBlock Block
	blocks       []Block
}

func (bc *Blockchain) newBlockchain(genesisBlock Block) {
	bc.genesisBlock = genesisBlock
	bc.blocks = []Block{genesisBlock}
}

func (bc *Blockchain) addBlock(block Block) error {
	if err := bc.isValidBlock(block); err != nil {
		fmt.Printf(err.Error())
		return err
	}

	bc.blocks = append(bc.blocks, block)

	return nil
}

func (bc *Blockchain) isValidBlock(block Block) error {
	if bc.isInvalidIndex(block) {
		return MismatchedIndex{expectedIndex: len(bc.blocks), actualIndex: block.index}.doError()
	}
	if bc.isAchronTimestamp(block) {
		return AchronologicalTimestamp{expectedTimestamp: bc.blocks[len(bc.blocks)-1].unixTimeStamp, actualTimestamp: block.unixTimeStamp}.doError()
	}
	if bc.isInvalidTimestamp(block) {
		return InvalidTimestamp{expectedTimestamp: time.Now().UTC().UnixNano(), actualTimestamp: block.unixTimeStamp}.doError()
	}
	if bc.isInvalidPreviousHash(block) {
		return InvalidPreviousHash{expectedHash: bc.blocks[len(bc.blocks)-1].hash, actualHash: block.previousHash}.doError()
	}
	if bc.isInvalidGenesisBlock(block) {
		return InvalidGenesisBlock{expectedFormat: make([]byte, 32), actualFormat: bc.genesisBlock.hash}.doError()
	}

	for _, transaction := range block.transaction {
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
	return block.index != len(bc.blocks)
}

func (bc *Blockchain) isAchronTimestamp(block Block) bool {
	return block.unixTimeStamp < bc.blocks[len(bc.blocks)-1].unixTimeStamp
}

func (bc *Blockchain) isInvalidTimestamp(block Block) bool {
	return block.unixTimeStamp > time.Now().UTC().UnixNano()
}

func (bc *Blockchain) isInvalidPreviousHash(block Block) bool {
	return !bytes.Equal(block.previousHash, bc.blocks[len(bc.blocks)-1].hash)
}

func (bc *Blockchain) isInvalidGenesisBlock(block Block) bool {
	return !bytes.Equal(bc.genesisBlock.hash, make([]byte, 32))
}

func (bc *Blockchain) isInsufficientInputValue(transaction Transaction) error {
	if transaction.inputValue() < transaction.outputValue() {
		return InsufficientInputValue{expectedValue: transaction.outputValue(), actualValue: transaction.inputValue()}.doError()
	}

	return nil
}

func (bc *Blockchain) isInputValueLessZero(transaction Transaction) error {
	for _, input := range transaction.inputs {
		if input.value < 0 {
			return InvalidInput{expectedValue: 0, actualValue: input.value}.doError()
		}
	}

	return nil
}
