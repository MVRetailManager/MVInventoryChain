package blockchain

import (
	"bytes"
	"runtime"

	logging "github.com/MVRetailManager/MVInventoryChain/logging"
	"github.com/MVRetailManager/MVInventoryChain/wallet"
)

type TxOutput struct {
	Value         int
	PublicKeyHash []byte
}

type TxInput struct {
	ID          []byte
	OutputIndex int
	Signature   []byte
	PublicKey   []byte
}

func (in *TxInput) UsesKey(publicKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PublicKey)

	return bytes.Compare(lockingHash, publicKeyHash) == 0
}

func (out *TxOutput) Lock(address []byte) {
	finalKey, err := wallet.Base58Decode(address)
	if err != nil {
		logging.ErrorLogger.Printf("%v", err)
		runtime.Goexit()
	}

	out.PublicKeyHash = finalKey[1 : len(finalKey)-4]
}

func (out *TxOutput) IsLockedWithkey(publicKeyHash []byte) bool {
	return bytes.Compare(out.PublicKeyHash, publicKeyHash) == 0
}

func NewTxOutput(value int, address string) *TxOutput {
	txOutput := &TxOutput{value, nil}
	txOutput.Lock([]byte(address))

	return txOutput
}
