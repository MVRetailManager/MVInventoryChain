package main

import (
	"testing"
	"time"
)

// func BenchmarkTest(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		test()
// 	}
// }

// func ExampleMain() {
// 	main()
// 	// Output: Hello, world.
// }

// MismatchedIndex Tests
func TestNewBlockMismatchedIndexLess(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		-1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockMismatchedIndexGreater(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		2,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockMismatchedIndexEqual(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		0,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockMismatchedIndex(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// AchronologicalTimeStamp Test
func TestNewBlockAchronologicalTimeStamp(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		10,
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockChronologicalTimeStamp(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InvalidTimestamp Test
func TestNewBlockInvalidTimestamp(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano()+1000000000000000000,
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockValidTimestamp(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InvalidPreviousHash Test
func TestNewBlockInvalidPreviousHashGenesis(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockInvalidPreviousHash(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	bc.addBlock(*block)

	block2 := newBlock(
		2,
		time.Now().UTC().UnixNano(),
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		1,
		[]Transaction{},
	)

	block2.mine()

	if bc.addBlock(*block2) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestValidPreviousHashGenesis(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

func TestValidPreviousHash(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	bc.addBlock(*block)

	block2 := newBlock(
		2,
		time.Now().UTC().UnixNano(),
		block.hash,
		1,
		[]Transaction{},
	)

	block2.mine()

	if bc.addBlock(*block2) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InvalidGenesisBlock Test
func TestNewBlockInvalidGenesisBlock(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          []byte{1, 2, 3, 4, 5, 6},
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockValidGenesisBlock(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
	bc.newBlockchain(genesisBlock)

	block := newBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.hash,
		1,
		[]Transaction{},
	)

	block.mine()

	if bc.addBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InsufficientInputValue Test
func TestNewBlockInsufficientInputValue(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
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
						value:   1,
					},
				},
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

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockSufficientInputValue(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
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
						value:   1,
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

	if bc.addBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InvalidInput Test
func TestNewBlockInvalidInput(t *testing.T) {
	genesisBlock := Block{
		index:         0,
		unixTimeStamp: time.Now().UTC().UnixNano(),
		hash:          make([]byte, 32),
		previousHash:  nil,
		nonce:         0,
		difficulty:    0,
		transaction:   nil,
	}

	bc := Blockchain{}
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
						value:   -1,
					},
				},
				outputs: []Output{
					{
						index:   0,
						address: "Alice",
						value:   -1,
					},
				},
			},
		},
	)

	block.mine()

	if bc.addBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}
