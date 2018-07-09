package FiniteAutomaton

type Automaton struct{
	start Node
}

func NewAutomaton(start Node) Automaton {
	return Automaton{start:start}
}

func (automaton Automaton) Walk(tokens []string) (state map[string]string) {
	node := automaton.start
	state = map[string]string{}
	for _, token := range tokens {
		node, state = node.traverse(token, state)
	}
}