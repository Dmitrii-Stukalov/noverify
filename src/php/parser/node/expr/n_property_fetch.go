package expr

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// PropertyFetch node
type PropertyFetch struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Variable     node.Node
	Property     node.Node
}

// NewPropertyFetch node constructor
func NewPropertyFetch(Variable node.Node, Property node.Node) *PropertyFetch {
	return &PropertyFetch{
		FreeFloating: nil,
		Variable:     Variable,
		Property:     Property,
	}
}

// SetPosition sets node position
func (n *PropertyFetch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *PropertyFetch) GetPosition() *position.Position {
	return n.Position
}

func (n *PropertyFetch) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PropertyFetch) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Variable != nil {
		n.Variable.Walk(v)
	}

	if n.Property != nil {
		n.Property.Walk(v)
	}

	v.LeaveNode(n)
}
