package FiniteAutomaton

type DynamicNode struct {

}

func NewDynamicNode() {
	panic("implement me")
}

func (dynamicNode DynamicNode) Name() string {
	panic("implement me")
}

func (dynamicNode DynamicNode) Edges() []Node {
	panic("implement me")
}

func (dynamicNode DynamicNode) traverse(token string, currentState map[string]string) (next Node, futureState map[string]string) {
	panic("implement me")
}

