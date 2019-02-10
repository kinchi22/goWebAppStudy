package binarysearchtree

import (
	"math"
)

// Item is the type of binary search tree
type Item int

// Node is a single node in the tree
type Node struct {
	key   int
	value Item
	left  *Node
	right *Node
}

// Bst is the binary search tree
type Bst struct {
	root *Node
}

// Search searches the targe node using key from the tree
func (bst *Bst) Search(key int) *Node {
	return search(bst.root, key)
}

// search is the internal recursive function to find the node by key
// it returns nil if no node was found
func search(n *Node, k int) *Node {
	if n == nil || n.key == k {
		return n
	}
	if k < n.key {
		return search(n.left, k)
	}
	// k > n.key
	return search(n.right, k)
}

// Insert inserts the new node into the tree
func (bst *Bst) Insert(key int, value Item) {
	n := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insert(bst.root, n)
	}
}

// insert is the internal recursive function to insert new node
// by finding proper position recursively
func insert(n *Node, new *Node) {
	if new.key < n.key {
		if n.left == nil {
			n.left = new
		} else {
			insert(n.left, new)
		}
	} else {
		if n.right == nil {
			n.right = new
		} else {
			insert(n.right, new)
		}
	}
}

// Delete deletes the node from the tree using key
func (bst *Bst) Delete(key int) {
	delete(bst.root, key)
}

// delete finds target node recursively and delete it from the tree
func delete(n *Node, k int) *Node {
	// if the current node is not the target node
	if n == nil {
		return nil
	}
	if k < n.key {
		n.left = delete(n.left, k)
		return n
	}
	if k > n.key {
		n.right = delete(n.right, k)
		return n
	}

	// if the current node is the target node
	if n.left == nil && n.right == nil {
		// current node doesn't have any child node
		n = nil
		return nil
	} else if n.left == nil {
		// current node has right child node only
		n = n.right
		return n
	} else if n.right == nil {
		// current node has left child node only
		n = n.left
		return n
	} else {
		// current node has both left and right child
		n = max(n.right)
		delete(n.right, n.key)
		return n
	}
}

// max finds most right child
func max(n *Node) *Node {
	max := n
	for {
		if max != nil && max.right != nil {
			max = max.right
		} else {
			break
		}
	}
	return max
}

// min finds most left child
func min(n *Node) *Node {
	min := n
	for {
		if min != nil && min.left != nil {
			min = min.left
		} else {
			break
		}
	}
	return min
}

// Traverse visits all nodes in order
func (bst *Bst) Traverse(f func(Item)) {
	traverse(bst.root, f)
}

// traverse is the internal recursive function to visit the nodes
func traverse(n *Node, f func(Item)) {
	if n != nil {
		traverse(n.left, f)
		f(n.value)
		traverse(n.right, f)
	}
}

// Verify checks the binary tree whether it's
func (bst *Bst) Verify() bool {
	return isBST(bst.root, math.MinInt64, math.MaxInt64)
}

func isBST(n *Node, min, max int) bool {
	if n == nil {
		return true
	}
	if n.key < min || n.key > max {
		return false
	}

	return isBST(n.left, min, n.key-1) && isBST(n.right, n.key+1, max)
}
