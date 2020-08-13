package models

type Tree struct {
	parent *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (tree *Tree) SetParent(node *Node) {
	tree.parent = node
}
func (tree *Tree) SetParentValue(value float32) {
	node := NewNode(value)
	tree.SetParent(node)
}
func (tree *Tree) AppendValue(value float32) {
	node := NewNode(value)
	tree.parent.AppendNode(node)
}
func (tree *Tree) HasValue(value float32) bool {
	return tree.parent.hasValue(value)
}
