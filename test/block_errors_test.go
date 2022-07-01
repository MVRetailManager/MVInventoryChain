package test

import (
	"testing"
	"time"

	blockchainPKG "github.com/MVRetailManager/MVInventoryChain/blockchain"
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
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		-1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockMismatchedIndexGreater(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		2,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockMismatchedIndexEqual(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		0,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockMismatchedIndex(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// AchronologicalTimeStamp Test
func TestNewBlockAchronologicalTimeStamp(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		10,
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockChronologicalTimeStamp(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InvalidTimestamp Test
func TestNewBlockInvalidTimestamp(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano()+1000000000000000000,
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockValidTimestamp(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InvalidPreviousHash Test
func TestNewBlockInvalidPreviousHashGenesis(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockInvalidPreviousHash(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) != nil {
		t.Errorf("Unexpected error")
	}

	block2 := blockchainPKG.NewBlock(
		2,
		time.Now().UTC().UnixNano(),
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		1,
		[]blockchainPKG.Transaction{},
	)

	block2.Mine()

	if bc.AddBlock(*block2) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestValidPreviousHashGenesis(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

func TestValidPreviousHash(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) != nil {
		t.Errorf("Unexpected error")
	}

	block2 := blockchainPKG.NewBlock(
		2,
		time.Now().UTC().UnixNano(),
		block.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block2.Mine()

	if bc.AddBlock(*block2) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InvalidGenesisBlock Test
func TestNewBlockInvalidGenesisBlock(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          []byte{1, 2, 3, 4, 5, 6},
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockValidGenesisBlock(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

	block := blockchainPKG.NewBlock(
		1,
		time.Now().UTC().UnixNano(),
		genesisBlock.Hash,
		1,
		[]blockchainPKG.Transaction{},
	)

	block.Mine()

	if bc.AddBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InsufficientInputValue Test
func TestNewBlockInsufficientInputValue(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

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
						Address: "Bob",
						Value:   1,
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

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}

func TestNewBlockSufficientInputValue(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

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
						Address: "Bob",
						Value:   1,
					},
				},
				Outputs: []blockchainPKG.Output{
					{
						Index:   0,
						Address: "Alice",
						Value:   1,
					},
				},
			},
		},
	)

	block.Mine()

	if bc.AddBlock(*block) != nil {
		t.Errorf("Expected false, got true")
	}
}

// InvalidInput Test
func TestNewBlockInvalidInput(t *testing.T) {
	genesisBlock := blockchainPKG.Block{
		Index:         0,
		UnixTimeStamp: time.Now().UTC().UnixNano(),
		Hash:          make([]byte, 32),
		PreviousHash:  nil,
		Nonce:         0,
		Difficulty:    0,
		Transaction:   nil,
	}

	bc := blockchainPKG.Blockchain{}
	bc.NewBlockchain(genesisBlock)

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
						Address: "Bob",
						Value:   -1,
					},
				},
				Outputs: []blockchainPKG.Output{
					{
						Index:   0,
						Address: "Alice",
						Value:   -1,
					},
				},
			},
		},
	)

	block.Mine()

	if bc.AddBlock(*block) == nil {
		t.Errorf("Expected false, got true")
	}
}
