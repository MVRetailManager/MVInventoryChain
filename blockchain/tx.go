package blockchain

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID          []byte
	OutputIndex int
	Sig         string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(address string) bool {
	return out.PubKey == address
}
