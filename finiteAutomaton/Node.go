package finiteAutomaton

type Node interface {
	Name() string
	Edges() []Node
	traverse(token string, currentState map[string]string) (next Node, futureState map[string]string)
}
