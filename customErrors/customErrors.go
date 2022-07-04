package customErrors

/*
import (
	"errors"
	"fmt"
)

// MismatchedIndex is an error type that is returned when the index of a block
// does not match the index of the previous block.
type MismatchedIndex struct {
	ExpectedIndex int
	ActualIndex   int
	statusCode    int
	err           error
}

func (mi *MismatchedIndex) Error() string {
	return fmt.Sprintf("MismatchedIndexError: Expected: %d, Got: %d", mi.ExpectedIndex, mi.ActualIndex)
}

func (mi MismatchedIndex) DoError() error {
	return &MismatchedIndex{
		ExpectedIndex: mi.ExpectedIndex,
		ActualIndex:   mi.ActualIndex,
		statusCode:    503,
		err:           errors.New("MismatchedIndexError"),
	}
}

// AchronologicalTimestamp is an error type that is returned when the timestamp
// of a block is earlier than the timestamp of the previous block.
type AchronologicalTimestamp struct {
	ExpectedTimestamp int64
	ActualTimestamp   int64
	statusCode        int
	err               error
}

func (ai *AchronologicalTimestamp) Error() string {
	return fmt.Sprintf("AchronologicalTimestampError: Expected: Timestamp > %d, Got: %d", ai.ExpectedTimestamp, ai.ActualTimestamp)
}

func (ai AchronologicalTimestamp) DoError() error {
	return &AchronologicalTimestamp{
		ExpectedTimestamp: ai.ExpectedTimestamp,
		ActualTimestamp:   ai.ActualTimestamp,
		statusCode:        503,
		err:               errors.New("AchronologicalTimestampError"),
	}
}

// InvalidTimestamp is an error type that is returned when the timestamp of a
// block has not happened yet.
type InvalidTimestamp struct {
	ExpectedTimestamp int64
	ActualTimestamp   int64
	statusCode        int
	err               error
}

func (it *InvalidTimestamp) Error() string {
	return fmt.Sprintf("InvalidTimestampError: Expected: Timestamp < %d, Got: %d", it.ExpectedTimestamp, it.ActualTimestamp)
}

func (it InvalidTimestamp) DoError() error {
	return &InvalidTimestamp{
		ExpectedTimestamp: it.ExpectedTimestamp,
		ActualTimestamp:   it.ActualTimestamp,
		statusCode:        503,
		err:               errors.New("InvalidTimestampError"),
	}
}

// InvalidPreviousHash is an error type that is returned when the previous hash
// of a block does not match the hash of the previous block.
type InvalidPreviousHash struct {
	ExpectedHash []byte
	ActualHash   []byte
	statusCode   int
	err          error
}

func (ip *InvalidPreviousHash) Error() string {
	return fmt.Sprintf("InvalidPreviousHashError: Expected: %x, Got: %x", ip.ExpectedHash, ip.ActualHash)
}

func (ip InvalidPreviousHash) DoError() error {
	return &InvalidPreviousHash{
		ExpectedHash: ip.ExpectedHash,
		ActualHash:   ip.ActualHash,
		statusCode:   503,
		err:          errors.New("InvalidPreviousHashError"),
	}
}

// InvalidGenesisBlock is an error type that is returned when the genesis block
// is invalid.
type InvalidGenesisBlock struct {
	ExpectedFormat []byte
	ActualFormat   []byte
	statusCode     int
	err            error
}

func (ig *InvalidGenesisBlock) Error() string {
	return fmt.Sprintf("InvalidGenesisBlockError: Expected format: %x, Got: %x", ig.ExpectedFormat, ig.ActualFormat)
}

func (ig InvalidGenesisBlock) DoError() error {
	return &InvalidGenesisBlock{
		ExpectedFormat: ig.ExpectedFormat,
		ActualFormat:   ig.ActualFormat,
		statusCode:     503,
		err:            errors.New("InvalidGenesisBlockError"),
	}
}

// InsufficientInputValue is an error type that is returned when the input value
// of a transaction is less than the minimum value.
type InsufficientInputValue struct {
	ExpectedValue int
	ActualValue   int
	statusCode    int
	err           error
}

func (ii *InsufficientInputValue) Error() string {
	return fmt.Sprintf("InsufficientInputValueError: Expected: <= %d, Got: %d", ii.ExpectedValue, ii.ActualValue)
}

func (ii InsufficientInputValue) DoError() error {
	return &InsufficientInputValue{
		ExpectedValue: ii.ExpectedValue,
		ActualValue:   ii.ActualValue,
		statusCode:    503,
		err:           errors.New("InsufficientInputValueError"),
	}
}

// InvalidInput is an error type that is returned when the input of a transaction
// is < 0.
type InvalidInput struct {
	ExpectedValue int
	ActualValue   int
	statusCode    int
	err           error
}

func (ii *InvalidInput) Error() string {
	return fmt.Sprintf("InvalidInputError: Expected: > %d, Got: %d", ii.ExpectedValue, ii.ActualValue)
}

func (ii InvalidInput) DoError() error {
	return &InvalidInput{
		ExpectedValue: ii.ExpectedValue,
		ActualValue:   ii.ActualValue,
		statusCode:    503,
		err:           errors.New("InvalidInputError"),
	}
}
*/
