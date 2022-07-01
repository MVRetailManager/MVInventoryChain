package blockchain

import (
	"github.com/dgraph-io/badger"

	logging "github.com/MVRetailManager/MVInventoryChain/logging"
)

const (
	dbPath = "./tmp/blockchain"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockchainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (bc *Blockchain) NewBlockchain(genesisBlock Block) {
	opts := badger.Options{}
	opts = badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	HandleError(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			err := txn.Set(genesisBlock.Hash, genesisBlock.Serialize())
			HandleError(err)
			err = txn.Set([]byte("lh"), genesisBlock.Hash)

			if err != nil {
				logging.ErrorLogger.Printf(err.Error())
				return err
			}

			bc.LastHash = genesisBlock.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			HandleError(err)
			bc.LastHash, err = item.ValueCopy(nil)
			return err
		}
	})

	HandleError(err)

	bc.Database = db

	logging.InfoLogger.Printf("New blockchain created with genesis block: %s", genesisBlock.Hash)
}

func (bc *Blockchain) AddBlock(block Block) error {
	/*if err := bc.isValidBlock(block); err != nil {
		logging.WarningLogger.Printf(err.Error())
		return err
	}*/

	err := bc.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		HandleError(err)
		bc.LastHash, err = item.ValueCopy(nil)

		HandleError(err)
		return err
	})

	HandleError(err)

	err = bc.Database.Update(func(txn *badger.Txn) error {
		err = txn.Set(block.Hash, block.Serialize())
		HandleError(err)

		err = txn.Set([]byte("lh"), block.Hash)

		if err != nil {
			logging.ErrorLogger.Printf(err.Error())
			return err
		}

		bc.LastHash = block.Hash

		return err
	})

	HandleError(err)

	return nil
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	iter := &BlockchainIterator{bc.LastHash, bc.Database}

	return iter
}

func (iter *BlockchainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		HandleError(err)

		encodedBlock, err := item.ValueCopy(nil)
		block = Deserialize(encodedBlock)

		return err
	})
	HandleError(err)

	iter.CurrentHash = block.PreviousHash

	return block
}

/*func (bc *Blockchain) isValidBlock(block Block) error {
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
}*/

func HandleError(err error) {
	if err != nil {
		logging.ErrorLogger.Printf(err.Error())
	}
}
