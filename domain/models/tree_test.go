package models

import "testing"

func TestTreeCreation(t *testing.T) {
	tree := NewTree()
	if tree.parent != nil {
		t.Error("The parent of the tree must be nil !")
	}
}

func TestSetTreeParent(t *testing.T) {
	tree := NewTree()
	node := NewNode(0)
	tree.SetParent(node)
	if tree.parent != node {
		t.Error("The parent of the tree must be the node !")
	}
}
func TestSetTreeParentByValue(t *testing.T) {
	tree := NewTree()
	tree.SetParentByValue(0)
	if tree.parent == nil {
		t.Error("The parent of the tree must not be nil !")
	}
	if tree.parent.value != 0 {
		t.Error("The value of the parent of the tree must be 0 !")
	}

}
