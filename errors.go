package main

import (
	"errors"
	"fmt"
)

// MismatchedIndex is an error type that is returned when the index of a block
// does not match the index of the previous block.
type MismatchedIndex struct {
	expectedIndex int
	actualIndex   int
	statusCode    int
	err           error
}

func (mi *MismatchedIndex) Error() string {
	return fmt.Sprintf("MismatchedIndexError: Expected: %d, Got: %d", mi.expectedIndex, mi.actualIndex)
}

func (mi MismatchedIndex) doError() error {
	return &MismatchedIndex{
		expectedIndex: mi.expectedIndex,
		actualIndex:   mi.actualIndex,
		statusCode:    503,
		err:           errors.New("MismatchedIndexError"),
	}
}

// AchronologicalTimestamp is an error type that is returned when the timestamp
// of a block is earlier than the timestamp of the previous block.
type AchronologicalTimestamp struct {
	expectedTimestamp int64
	actualTimestamp   int64
	statusCode        int
	err               error
}

func (ai *AchronologicalTimestamp) Error() string {
	return fmt.Sprintf("AchronologicalTimestampError: Expected: Timestamp > %d, Got: %d", ai.expectedTimestamp, ai.actualTimestamp)
}

func (ai AchronologicalTimestamp) doError() error {
	return &AchronologicalTimestamp{
		expectedTimestamp: ai.expectedTimestamp,
		actualTimestamp:   ai.actualTimestamp,
		statusCode:        503,
		err:               errors.New("AchronologicalTimestampError"),
	}
}

// InvalidTimestamp is an error type that is returned when the timestamp of a
// block has not happened yet.
type InvalidTimestamp struct {
	expectedTimestamp int64
	actualTimestamp   int64
	statusCode        int
	err               error
}

func (it *InvalidTimestamp) Error() string {
	return fmt.Sprintf("InvalidTimestampError: Expected: Timestamp < %d, Got: %d", it.expectedTimestamp, it.actualTimestamp)
}

func (it InvalidTimestamp) doError() error {
	return &InvalidTimestamp{
		expectedTimestamp: it.expectedTimestamp,
		actualTimestamp:   it.actualTimestamp,
		statusCode:        503,
		err:               errors.New("InvalidTimestampError"),
	}
}

// InvalidPreviousHash is an error type that is returned when the previous hash
// of a block does not match the hash of the previous block.
type InvalidPreviousHash struct {
	expectedHash []byte
	actualHash   []byte
	statusCode   int
	err          error
}

func (ip *InvalidPreviousHash) Error() string {
	return fmt.Sprintf("InvalidPreviousHashError: Expected: %x, Got: %x", ip.expectedHash, ip.actualHash)
}

func (ip InvalidPreviousHash) doError() error {
	return &InvalidPreviousHash{
		expectedHash: ip.expectedHash,
		actualHash:   ip.actualHash,
		statusCode:   503,
		err:          errors.New("InvalidPreviousHashError"),
	}
}

// InvalidGenesisBlock is an error type that is returned when the genesis block
// is invalid.
type InvalidGenesisBlock struct {
	expectedFormat []byte
	actualFormat   []byte
	statusCode     int
	err            error
}

func (ig *InvalidGenesisBlock) Error() string {
	return fmt.Sprintf("InvalidGenesisBlockError: Expected format: %x, Got: %x", ig.expectedFormat, ig.actualFormat)
}

func (ig InvalidGenesisBlock) doError() error {
	return &InvalidGenesisBlock{
		expectedFormat: ig.expectedFormat,
		actualFormat:   ig.actualFormat,
		statusCode:     503,
		err:            errors.New("InvalidGenesisBlockError"),
	}
}

// InsufficientInputValue is an error type that is returned when the input value
// of a transaction is less than the minimum value.
type InsufficientInputValue struct {
	expectedValue int
	actualValue   int
	statusCode    int
	err           error
}

func (ii *InsufficientInputValue) Error() string {
	return fmt.Sprintf("InsufficientInputValueError: Expected: > %d, Got: %d", ii.expectedValue, ii.actualValue)
}

func (ii InsufficientInputValue) doError() error {
	return &InsufficientInputValue{
		expectedValue: ii.expectedValue,
		actualValue:   ii.actualValue,
		statusCode:    503,
		err:           errors.New("InsufficientInputValueError"),
	}
}

// InvalidInput is an error type that is returned when the input of a transaction
// is < 0.
type InvalidInput struct {
	expectedValue int
	actualValue   int
	statusCode    int
	err           error
}

func (ii *InvalidInput) Error() string {
	return fmt.Sprintf("InvalidInputError: Expected: > %d, Got: %d", ii.expectedValue, ii.actualValue)
}

func (ii InvalidInput) doError() error {
	return &InvalidInput{
		expectedValue: ii.expectedValue,
		actualValue:   ii.actualValue,
		statusCode:    503,
		err:           errors.New("InvalidInputError"),
	}
}
