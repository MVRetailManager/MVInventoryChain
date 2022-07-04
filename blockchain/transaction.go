package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"math/big"
	"runtime"
	"strings"

	"github.com/MVRetailManager/MVInventoryChain/logging"
	"github.com/MVRetailManager/MVInventoryChain/wallet"
)

const (
	reward = 100
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	HandleError(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	twin := TxInput{[]byte{}, -1, nil, []byte(data)}
	txout := NewTxOutput(reward, to)

	tx := Transaction{nil, []TxInput{twin}, []TxOutput{*txout}}
	tx.SetID()

	return &tx
}

func NewTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	wallets, err := wallet.CreateWallets()
	HandleError(err)

	w := wallets.GetWallet(from)
	publicKeyHash := wallet.PublicKeyHash(w.PublicKey)

	acc, validOutputs := bc.FindSpendableOutputs(publicKeyHash, amount)

	if acc < amount {
		fmt.Printf("Not enough funds on address %s\n", from)
		logging.WarningLogger.Printf("Not enough funds on address %s\n", from)
		runtime.Goexit()
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		HandleError(err)

		for _, out := range outs {
			input := TxInput{txID, out, nil, w.PublicKey}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, *NewTxOutput(amount, to))

	if acc > amount {
		outputs = append(outputs, *NewTxOutput(acc-amount, from))
	}

	tx := Transaction{nil, inputs, outputs}
	tx.ID = tx.Hash()

	bc.SignTransaction(&tx, w.PrivateKey)

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].OutputIndex == -1
}

func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	HandleError(err)

	return encoded.Bytes()
}

func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	return hash[:]
}

func (tx *Transaction) Sign(privateKey ecdsa.PrivateKey, previousTxs map[string]Transaction) {
	if tx.IsCoinbase() {
		return
	}

	for _, in := range tx.Inputs {
		if previousTxs[hex.EncodeToString(in.ID)].ID == nil {
			logging.ErrorLogger.Printf("Error with transaction %s, previous transaction does not exist.", hex.EncodeToString(in.ID))
			runtime.Goexit()
		}
	}

	txCopy := tx.TrimmedCopy()

	for inId, in := range txCopy.Inputs {
		prevTx := previousTxs[hex.EncodeToString(in.ID)]

		txCopy.Inputs[inId].Signature = nil
		txCopy.Inputs[inId].PublicKey = prevTx.Outputs[in.OutputIndex].PublicKeyHash

		txCopy.ID = txCopy.Hash()

		txCopy.Inputs[inId].PublicKey = nil

		r, s, err := ecdsa.Sign(rand.Reader, &privateKey, txCopy.ID)
		HandleError(err)

		signature := append(r.Bytes(), s.Bytes()...)
		tx.Inputs[inId].Signature = signature
	}
}

func (tx Transaction) TrimmedCopy() Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	for _, in := range tx.Inputs {
		inputs = append(inputs, TxInput{in.ID, in.OutputIndex, nil, nil})
	}

	for _, out := range tx.Outputs {
		outputs = append(outputs, TxOutput{out.Value, out.PublicKeyHash})
	}

	txCopy := Transaction{tx.ID, inputs, outputs}

	return txCopy
}

func (tx *Transaction) Verify(previousTxs map[string]Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	for _, in := range tx.Inputs {
		if previousTxs[hex.EncodeToString(in.ID)].ID == nil {
			logging.ErrorLogger.Printf("Error with transaction %s, previous transaction does not exist.", hex.EncodeToString(in.ID))
			runtime.Goexit()
		}
	}

	txCopy := tx.TrimmedCopy()
	curve := elliptic.P256()

	for inId, in := range txCopy.Inputs {
		prevTx := previousTxs[hex.EncodeToString(in.ID)]

		txCopy.Inputs[inId].Signature = nil
		txCopy.Inputs[inId].PublicKey = prevTx.Outputs[in.OutputIndex].PublicKeyHash

		txCopy.ID = txCopy.Hash()

		txCopy.Inputs[inId].PublicKey = nil

		r := big.Int{}
		s := big.Int{}

		sigLen := len(in.Signature)

		r.SetBytes(in.Signature[:(sigLen / 2)])
		s.SetBytes(in.Signature[(sigLen / 2):])

		x := big.Int{}
		y := big.Int{}

		keyLen := len(in.PublicKey)

		x.SetBytes(in.PublicKey[:(keyLen / 2)])
		y.SetBytes(in.PublicKey[(keyLen / 2):])

		rawPublicKey := ecdsa.PublicKey{curve, &x, &y}

		if ecdsa.Verify(&rawPublicKey, txCopy.ID, &r, &s) == false {
			return false
		}
	}

	return true
}

func (tx Transaction) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("--- Transaction %x:", tx.ID))

	for i, input := range tx.Inputs {
		lines = append(lines, fmt.Sprintf("		Input 			%d", i))
		lines = append(lines, fmt.Sprintf("		TXID: 			%d", input.ID))
		lines = append(lines, fmt.Sprintf("		OutputIndex: 		%d", input.OutputIndex))
		lines = append(lines, fmt.Sprintf("		Signature: 		%x", input.Signature))
		lines = append(lines, fmt.Sprintf("		PublicKey: 		%x", input.PublicKey))
	}

	for i, output := range tx.Outputs {
		lines = append(lines, fmt.Sprintf("		Output 			%d", i))
		lines = append(lines, fmt.Sprintf("		Value: 			%d", output.Value))
		lines = append(lines, fmt.Sprintf("		PublicKeyHash: 	%x", output.PublicKeyHash))
	}

	return strings.Join(lines, "\n")
}
