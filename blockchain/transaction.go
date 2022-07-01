package blockchain

type Transaction struct {
	Inputs  []Output
	Outputs []Output
}

func (t *Transaction) inputValue() int {
	var total int

	for _, input := range t.Inputs {
		total += input.Value
	}

	return total
}

func (t *Transaction) outputValue() int {
	var total int

	for _, output := range t.Outputs {
		total += output.Value
	}

	return total
}
