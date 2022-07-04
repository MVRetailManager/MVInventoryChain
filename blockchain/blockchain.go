package blockchain

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/dgraph-io/badger"

	logging "github.com/MVRetailManager/MVInventoryChain/logging"
)

const (
	dbPath      = "./tmp/blockchain"
	dbFile      = "./tmp/blockchain/MANIFEST"
	genesisData = "Genesis Block"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockchainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (bc *Blockchain) InitBlockchain(address string) {
	var genesis Block

	if DBexists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	opts := InitDBOpts()

	db, err := badger.Open(opts)
	HandleError(err)

	err = db.Update(func(txn *badger.Txn) error {
		cbtx := CoinbaseTx(address, genesisData)
		genesis = initGenesis([]*Transaction{cbtx})
		err = txn.Set(genesis.Hash, genesis.Serialize())
		HandleError(err)
		err = txn.Set([]byte("lh"), genesis.Hash)

		bc.LastHash = genesis.Hash

		return err
	})

	HandleError(err)

	bc.Database = db

	logging.InfoLogger.Printf("New blockchain created with genesis block: %s", genesis.Hash)
}

func DBexists() bool {
	_, err := os.Stat(dbFile)

	return err == nil
}

func (bc *Blockchain) ContinueBlockchain(address string) {
	if !DBexists() {
		fmt.Println("No existing blockchain found, please create one.")
		runtime.Goexit()
	}

	opts := InitDBOpts()

	db, err := badger.Open(opts)
	HandleError(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		HandleError(err)

		err = item.Value(func(val []byte) error {
			bc.LastHash = val
			return nil
		})

		return err
	})
	HandleError(err)

	bc.Database = db
}

func (bc *Blockchain) FindUnspentTxs(address string) []Transaction {
	var unspentTxs []Transaction

	spentTXOs := make(map[string][]int)

	iter := bc.Iterator()

	for {
		block, err := iter.Next()
		HandleError(err)

		for _, tx := range block.Transaction {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Outputs {
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				if out.CanBeUnlocked(address) {
					unspentTxs = append(unspentTxs, *tx)
				}
			}
			if !tx.IsCoinbase() {
				for _, in := range tx.Inputs {
					if in.CanUnlock(address) {
						inTxID := hex.EncodeToString(in.ID)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.OutputIndex)
					}
				}
			}
		}

		if len(block.PreviousHash) == 0 {
			break
		}
	}

	return unspentTxs
}

func (bc *Blockchain) HandleUnspentTxs(address string) []TxOutput {
	var unspentTxs []TxOutput

	unspentTransactions := bc.FindUnspentTxs(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Outputs {
			if out.CanBeUnlocked(address) {
				unspentTxs = append(unspentTxs, out)
			}
		}
	}

	return unspentTxs
}

func (bc *Blockchain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	unspentTxs := bc.FindUnspentTxs(address)
	acc := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Outputs {
			if out.CanBeUnlocked(address) && acc < amount {
				acc = out.Value
				unspentOuts[txID] = append(unspentOuts[txID], outIdx)

				if acc >= amount {
					break Work
				}
			}
		}
	}

	return acc, unspentOuts
}

func InitDBOpts() badger.Options {
	opts := badger.Options{}
	opts = badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath
	opts.Logger = nil

	return opts
}

func (bc *Blockchain) AddBlock(block Block) {
	err := bc.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		HandleError(err)
		err = item.Value(func(val []byte) error {
			bc.LastHash = val
			return nil
		})

		return err
	})
	HandleError(err)

	nbIndex, _ := bc.Database.Size()
	newBlock := NewBlock(int(nbIndex), time.Now().UnixNano(), bc.LastHash, block.Transaction)

	err = bc.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		HandleError(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		bc.LastHash = newBlock.Hash

		return err
	})

	HandleError(err)
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	iter := &BlockchainIterator{bc.LastHash, bc.Database}

	return iter
}

func (iter *BlockchainIterator) Next() (*Block, error) {
	var block *Block
	var encodedBlock []byte

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)

		if err != nil {
			HandleError(err)
			return err
		}

		err = item.Value(func(val []byte) error {
			encodedBlock = val
			return nil
		})

		block = Deserialize(encodedBlock)

		return err
	})

	if err != nil {
		HandleError(err)
		return nil, err
	}

	iter.CurrentHash = block.PreviousHash

	return block, nil
}

func initGenesis(coinbase []*Transaction) Block {
	return Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   coinbase,
	}
}

func HandleError(err error) {
	if err != nil {
		logging.ErrorLogger.Printf(err.Error())
	}
}
