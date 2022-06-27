package main

import (
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"strings"
)

type Block struct {
	index         int
	unixTimeStamp int64
	hash          []byte
	previousHash  []byte
	nonce         int
	difficulty    int
	transaction   []Transaction
}

func newBlock(index int, unixTimeStamp int64, previousHash []byte, difficulty int, transaction []Transaction) *Block {
	b := &Block{
		index:         index,
		unixTimeStamp: unixTimeStamp,
		hash:          make([]byte, 32),
		previousHash:  previousHash,
		nonce:         0,
		difficulty:    difficulty,
		transaction:   transaction,
	}

	b.mine()

	return b
}

func (b *Block) mine() {
	for !b.validateHash(b.calculateHash()) {
		b.nonce++
		b.hash = b.calculateHash()
	}
}

func (b *Block) calculateHash() []byte {
	hash := sha256.Sum256([]byte(strconv.Itoa(b.index) + strconv.FormatInt(b.unixTimeStamp, 10) + string(b.previousHash) + strconv.Itoa(b.nonce) + strconv.Itoa(b.difficulty) + transactionToString(b.transaction)))
	return hash[:]
}

func (b *Block) validateHash(hash []byte) bool {
	prefix := strings.Repeat("0", b.difficulty)
	return strings.HasPrefix(string(hash), prefix)
}

func first(n []byte, _ error) []byte {
	return n
}

func transactionToString(transaction []Transaction) string {
	return string(first(json.Marshal(transaction)))
}
