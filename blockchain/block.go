package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"strings"

	logging "github.com/MVRetailManager/MVInventoryChain/logging"
)

type Block struct {
	Index         int
	UnixTimeStamp int64
	Hash          []byte
	PreviousHash  []byte
	Nonce         int
	Difficulty    int
	Transaction   []Transaction
}

func NewBlock(index int, unixTimeStamp int64, previousHash []byte, difficulty int, transaction []Transaction) *Block {
	b := &Block{
		Index:         index,
		UnixTimeStamp: unixTimeStamp,
		Hash:          make([]byte, 32),
		PreviousHash:  previousHash,
		Nonce:         0,
		Difficulty:    difficulty,
		Transaction:   transaction,
	}

	logging.BlocksLogger.Printf("Block Initialised: %v\n", b)

	b.Mine()

	return b
}

func (b *Block) Mine() {
	for !b.validateHash(b.calculateHash()) {
		b.Nonce++
		b.Hash = b.calculateHash()
	}

	logging.BlocksLogger.Printf("Block Mined: %v\n", b)
}

func (b *Block) calculateHash() []byte {
	hash := sha256.Sum256([]byte(strconv.Itoa(b.Index) + strconv.FormatInt(b.UnixTimeStamp, 10) + string(b.PreviousHash) + strconv.Itoa(b.Nonce) + strconv.Itoa(b.Difficulty) + transactionToString(b.Transaction)))
	return hash[:]
}

func (b *Block) validateHash(hash []byte) bool {
	prefix := strings.Repeat("0", b.Difficulty)
	return strings.HasPrefix(string(hash), prefix)
}

func first(n []byte, _ error) []byte {
	return n
}

func transactionToString(transaction []Transaction) string {
	return string(first(json.Marshal(transaction)))
}
