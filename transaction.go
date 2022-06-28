package main

type Transaction struct {
	inputs  []Output
	outputs []Output
}

func (t *Transaction) inputValue() int {
	var total int

	for _, input := range t.inputs {
		total += input.value
	}

	return total
}

func (t *Transaction) outputValue() int {
	var total int

	for _, output := range t.outputs {
		total += output.value
	}

	return total
}
