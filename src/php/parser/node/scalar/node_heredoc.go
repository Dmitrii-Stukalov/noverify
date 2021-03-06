package scalar

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// Heredoc node
type Heredoc struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Label        string
	Parts        []node.Node
}

// NewHeredoc node constructor
func NewHeredoc(Label string, Parts []node.Node) *Heredoc {
	return &Heredoc{
		FreeFloating: nil,
		Label:        Label,
		Parts:        Parts,
	}
}

// SetPosition sets node position
func (n *Heredoc) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Heredoc) GetPosition() *position.Position {
	return n.Position
}

func (n *Heredoc) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Heredoc) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Parts != nil {
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(v)
			}
		}
	}

	v.LeaveNode(n)
}
