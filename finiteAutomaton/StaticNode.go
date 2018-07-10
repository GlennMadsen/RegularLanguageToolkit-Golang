package finiteAutomaton

type StaticNode struct {
	name      string
	traverser func(string, map[string]string) (Node, map[string]string)
	edges     []Node
}

func NewStaticNode(name string, traverser func(string, map[string]string) (next Node, futureState map[string]string), edges []Node) StaticNode {
	return StaticNode{name: name, traverser: traverser, edges: edges}
}

func (staticNode StaticNode) Name() string {
	return staticNode.name
}

func (staticNode StaticNode) Edges() []Node {
	return staticNode.edges
}

func (staticNode StaticNode) traverse(token string, currentState map[string]string) (next Node, futureState map[string]string) {
	next, futureState = staticNode.traverser(token, currentState)
	return
}
