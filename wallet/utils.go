package wallet

import (
	"github.com/mr-tron/base58"
)

func Base58Encode(input []byte) []byte {
	return []byte(base58.Encode(input))
}

func Base58Decode(input []byte) ([]byte, error) {
	return base58.Decode(string(input[:]))
}
