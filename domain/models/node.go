package models

type Node struct {
	value     float32
	letfNode  *Node
	rightNode *Node
}

func NewNode(value float32) *Node {
	return &Node{value: value}
}
func (node *Node) AppendNode(child *Node) {
	if child.value > node.value {
		if node.rightNode == nil {
			node.rightNode = child
		} else {
			node.rightNode.AppendNode(child)
		}
	} else if child.value < node.value {
		if node.letfNode == nil {
			node.letfNode = child
		} else {
			node.letfNode.AppendNode(child)
		}
	}
}

func (node *Node) hasValue(value float32) bool {
	if node.value == value {
		return true
	} else {
		if value > node.value {
			if node.rightNode == nil {
				return false
			} else {
				return node.rightNode.hasValue(value)
			}
		} else {
			if node.letfNode == nil {
				return false
			} else {
				return node.letfNode.hasValue(value)
			}
		}
	}
}
