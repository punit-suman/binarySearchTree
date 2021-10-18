package binarySearchTree

import "fmt"

// In-order traversal
// Left branch -> Current Node -> Right Branch
// visits node in ascending order
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Println(n.Value)
	n.Right.InOrderTraversal()
}

// Pre-order traversal
// Current Node -> Left Node -> Right Node
// Root Node is visited first
func (n *Node) PreOrderTraversal() {
	if n == nil {
		return
	}
	fmt.Println(n.Value)
	n.Left.PreOrderTraversal()
	n.Right.PreOrderTraversal()
}

// Post-order traversal
// Left Node -> Right Node -> Current Node
// Root Node is visited last
func (n *Node) PostOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.PostOrderTraversal()
	n.Right.PostOrderTraversal()
	fmt.Println(n.Value)
}

func (n *Node) MinValueNode() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (n *Node) MaxValueNode() *Node {
	current := n
	for current.Right != nil {
		current = current.Right
	}
	return current
}

func (n *Node) LeftView(values ...int) int {
	var level, last int
	if len(values) > 0 {
		level = values[0]
		last = values[1]
	} else {
		level = 1
		last = 0
	}
	if n == nil {
		return last
	}
	if last < level {
		fmt.Println(n.Value)
		last = level
	}
	last = n.Left.LeftView(level+1, last)
	last = n.Right.LeftView(level+1, last)
	return last
}

func (n *Node) RightView(values ...int) int {
	var level, last int
	if len(values) > 0 {
		level = values[0]
		last = values[1]
	} else {
		level = 1
		last = 0
	}
	if n == nil {
		return last
	}
	if last < level {
		fmt.Println(n.Value)
		last = level
	}
	last = n.Right.RightView(level+1, last)
	last = n.Left.RightView(level+1, last)
	return last
}

func (n *Node) LeavesView() {
	if n == nil {
		return
	}
	if n.Left == nil && n.Right == nil {
		fmt.Println(n.Value)
	} else {
		n.Left.LeavesView()
		n.Right.LeavesView()
	}
}
