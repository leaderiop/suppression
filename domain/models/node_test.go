package models

import "testing"

func TestNodeCreation(t *testing.T) {
	var value float32
	node := NewNode(value)

	if node.value != value {
		t.Error("value isn't not correct! ")
	}
	if node.letfNode != nil {
		t.Error("left node must be nil !")
	}
	if node.rightNode != nil {
		t.Error("right node must be nil !")
	}
}
func TestAppendChild(t *testing.T) {
	node1 := NewNode(0)
	node2 := NewNode(2)
	node3 := NewNode(-1)
	node4 := NewNode(1)
	node1.AppendNode(node2)

	if node1.rightNode == nil {
		t.Error("The right child node of the node 1 must not be nil after appending node 2 !")
	}
	if node1.letfNode != nil {
		t.Error("The left child node of the node 1 must  be nil after appending node 2 !")
	}
	node1.AppendNode(node3)
	if node1.rightNode == nil {
		t.Error("The right child node of the node 1 must not be nil after appending node 3 !")
	}
	if node1.letfNode == nil {
		t.Error("The left child node of the node 1 must not be nil after appending node 3 !")
	}
	node1.AppendNode(node4)
	if node2.rightNode != nil {
		t.Error("The right child node of the node 2 must  be nil after appending node 4!")
	}
	if node2.letfNode == nil {
		t.Error("The left child node of the node 2 must  not be nil after appending node 4!")
	}

}
func TestHasValue(t *testing.T) {
	node1 := NewNode(0)
	node2 := NewNode(2)
	node3 := NewNode(-1)
	node4 := NewNode(1)
	node1.AppendNode(node2)
	node1.AppendNode(node3)
	node1.AppendNode(node4)

	if !node1.hasValue(0) {
		t.Error("The node 1 must have the value 0")
	}
	if !node1.hasValue(-1) {
		t.Error("The node 1 must have the value -1")
	}
	if !node1.hasValue(1) {
		t.Error("The node 1 must have the value 1")
	}
	if !node1.hasValue(2) {
		t.Error("The node 1 must have the value 2")
	}
	if node1.hasValue(3) {
		t.Error("The node 1 must not have the value 3")
	}
}
