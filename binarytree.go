package binarytree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Value < k {
		//move right if value is lesser
		if n.Right == nil {
			n.Right = &Node{Value: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Value > k {
		//move left if value is greater
		if n.Left == nil {
			n.Left = &Node{Value: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

var count int = 0

func (n *Node) NumberOfNodes() int {
	if n == nil {
		return count
	}
	if n != nil {
		n.Left.NumberOfNodes()
		n.Right.NumberOfNodes()
		count++
	}
	return count
}

var countLeaves int = 0

func (n *Node) NumberOfLeaves() int {
	if n == nil {
		return countLeaves
	}
	if n.Left == nil && n.Right == nil {
		countLeaves++
	} else {
		n.Left.NumberOfLeaves()
		n.Right.NumberOfLeaves()
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
