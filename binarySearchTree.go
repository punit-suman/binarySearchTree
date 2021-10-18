package binarySearchTree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	v := &Node{Value: k}
	if n.Value < k {
		//move right if value is lesser
		if n.Right == nil {
			n.Right = v
		} else {
			n.Right.Insert(k)
		}
	} else if n.Value >= k {
		//move left if value is greater or equal
		if n.Left == nil {
			n.Left = v
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) NumberOfNodes(c ...int) int {
	var count int
	if len(c) == 0 {
		count = 0
	} else {
		count = c[0]
	}
	if n == nil {
		return count
	}
	if n != nil {
		count++
		count = n.Left.NumberOfNodes(count)
		count = n.Right.NumberOfNodes(count)
	}
	return count
}

func (n *Node) NumberOfLeaves(l ...int) int {
	var countLeaves int
	if len(l) == 0 {
		countLeaves = 0
	} else {
		countLeaves = l[0]
	}
	if n == nil {
		return countLeaves
	}
	if n.Left == nil && n.Right == nil {
		countLeaves++
	} else {
		countLeaves = n.Left.NumberOfLeaves(countLeaves)
		countLeaves = n.Right.NumberOfLeaves(countLeaves)
	}
	return countLeaves
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Value < k {
		//move right
		return n.Right.Search(k)
	} else if n.Value > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}

func (n *Node) Delete(k int) *Node {
	var temp *Node
	if n == nil {
		return nil
	}
	if n.Value > k {
		n.Left = n.Left.Delete(k)
	} else if n.Value < k {
		n.Right = n.Right.Delete(k)
	} else {
		if n.Left == nil {
			temp = n.Right
			n = nil
			return temp
		} else if n.Right == nil {
			temp = n.Left
			n = nil
			return temp
		} else {
			temp = n.Right.MinValueNode()
			n.Value = temp.Value
			n.Right = n.Right.Delete(temp.Value)
		}
	}
	return n
}
