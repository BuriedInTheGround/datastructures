package binarysearchtree

import (
	"container/list"
	"fmt"
	"math"
)

// Node is a vertex of the BinarySearchTree.
type Node struct {
	data  int
	left  *Node
	right *Node
}

// BinarySearchTree is a binary tree data structure that satisfies the BST
// invariant: the left substree of every node has element with smaller value
// and the right subtree of every node has element with bigger value.
type BinarySearchTree struct {
	root *Node
	size int
}

// New returns a new BinarySearchTree instance.
func New() BinarySearchTree {
	return BinarySearchTree{root: nil, size: 0}
}

// Size returns the number of elements contained into the tree.
//
// Complexity: O(1)
func (bst *BinarySearchTree) Size() int {
	return bst.size
}

// IsEmpty returns whether the tree is empty or not.
//
// Complexity: O(1)
func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.Size() == 0
}

// TotalDegree returns the sum of the degree of every node of the tree.
//
// Complexity: O(1)
func (bst *BinarySearchTree) TotalDegree() int {
	if bst.IsEmpty() {
		panic("an empty tree does not have a degree")
	}
	return bst.Size() - 1
}

// Height returns the height of the tree.
//
// Complexity: O(n)
func (bst *BinarySearchTree) Height() int {
	return bst.height(bst.root)
}

func (bst *BinarySearchTree) height(node *Node) int {
	// A leaf has an height of zero.
	if node == nil {
		return 0
	}

	// The height of a node that is not a leaf is the maximum between the
	// height of the right subtree and the left subtree, plus one.
	leftHeight := float64(bst.height(node.left))
	rightHeight := float64(bst.height(node.right))
	return int(math.Max(leftHeight, rightHeight)) + 1
}

// Insert adds an node with the specified `value` into the tree, if it does
// not already exists, otherwise returns an error.
//
// Complexity: O(log(n)) on average, O(n) in the worst case
func (bst *BinarySearchTree) Insert(value int) error {
	err := bst.insert(&bst.root, value)
	if err != nil {
		return err
	}
	bst.size++
	return nil
}

func (bst *BinarySearchTree) insert(node **Node, value int) error {
	// If the pointed node is a leaf add a new node there.
	if *node == nil {
		*node = &Node{data: value, left: nil, right: nil}
		return nil
	}

	// Go down left or right depending on the value.
	if value < (*node).data {
		return bst.insert(&(*node).left, value)
	} else if value > (*node).data {
		return bst.insert(&(*node).right, value)
	}

	// If the value to be inserted is found, return an error: duplicate values
	// are not allowed.
	return fmt.Errorf("the value %d is already in the tree [%d]", value, (*node).data)
}

// Remove removes the node that contains the specified `value`, if exists, and
// restore the BST invariant.
//
// Complexity: O(log(n)) on average, O(n) in the worst case
func (bst *BinarySearchTree) Remove(value int) {
	// Do the removal only if the value exists inside the tree.
	if bst.Contains(value) {
		bst.root = *bst.remove(&bst.root, value)
		bst.size--
	}
}

func (bst *BinarySearchTree) remove(node **Node, value int) **Node {
	// If the pointed node is a leaf, return a leaf.
	if *node == nil {
		return nil
	}

	if value < (*node).data {
		// Remove the value from the left subtree.
		(*node).left = *bst.remove(&(*node).left, value)
	} else if value > (*node).data {
		// Remove the value from the right subtree.
		(*node).right = *bst.remove(&(*node).right, value)
	} else {
		// The node to remove is now `node`.

		// If the left subtree is empty, swap the node to remove with the
		// right subtree (even if it is empty).
		if (*node).left == nil {
			return &(*node).right
		}

		// If the right subtee is empty, swap the node to remove with the left
		// subtree (even if it is empty).
		if (*node).right == nil {
			return &(*node).left
		}

		// Otherwise both the left and the right subtrees exist.
		// So, take the smallest value of the right subtree, copy its data
		// into the node to remove, and finally remove the node from which the
		// data was copied to avoid duplicates.
		temp := bst.digLeft(&(*node).right)
		(*node).data = (*temp).data
		(*node).right = *bst.remove(&(*node).right, (*temp).data)

	}

	// Return the root of the removal.
	return node
}

func (bst *BinarySearchTree) digLeft(node **Node) **Node {
	for (*node).left != nil {
		node = &(*node).left
	}
	return node
}

func (bst *BinarySearchTree) digRight(node *Node) *Node {
	for node.right != nil {
		node = node.right
	}
	return node
}

// Contains returns whether the tree contains a node with the specified
// `value` or not.
//
// Complexity: O(log(n)) on average, O(n) in the worst case
func (bst *BinarySearchTree) Contains(value int) bool {
	if bst.IsEmpty() {
		return false
	}
	return bst.search(bst.root, value)
}

func (bst *BinarySearchTree) search(node *Node, value int) bool {
	// If the node is empty, it cannot contains any value, so return false.
	if node == nil {
		return false
	}

	// Go down left or right depending on the value.
	if value < node.data {
		return bst.search(node.left, value)
	} else if value > node.data {
		return bst.search(node.right, value)
	}

	// If the value was neither smaller nor greater, then it's found.
	return true
}

// TraversePreOrder traverses the tree nodes in a pre-order fashion, putting
// the values into a slice and returning it.
func (bst *BinarySearchTree) TraversePreOrder() []int {
	res := make([]int, 0, bst.Size())
	bst.preOrder(bst.root, &res)
	return res
}

func (bst *BinarySearchTree) preOrder(node *Node, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.data)
	bst.preOrder(node.left, res)
	bst.preOrder(node.right, res)
}

// TraverseInOrder traverses the tree nodes in an in-order fashion, putting
// the values into a slice and returning it.
func (bst *BinarySearchTree) TraverseInOrder() []int {
	res := make([]int, 0, bst.Size())
	bst.inOrder(bst.root, &res)
	return res
}

func (bst *BinarySearchTree) inOrder(node *Node, res *[]int) {
	if node == nil {
		return
	}
	bst.inOrder(node.left, res)
	*res = append(*res, node.data)
	bst.inOrder(node.right, res)
}

// TraversePostOrder traverses the tree nodes in a post-order fashion, putting
// the values into a slice and returning it.
func (bst *BinarySearchTree) TraversePostOrder() []int {
	res := make([]int, 0, bst.Size())
	bst.postOrder(bst.root, &res)
	return res
}

func (bst *BinarySearchTree) postOrder(node *Node, res *[]int) {
	if node == nil {
		return
	}
	bst.postOrder(node.left, res)
	bst.postOrder(node.right, res)
	*res = append(*res, node.data)
}

// TraverseLevelOrder traverses the tree nodes in a level-order fashion
// (basically doing a breadth first search), putting the values into a slice
// and returning it.
func (bst *BinarySearchTree) TraverseLevelOrder() []int {
	return bst.breadthFirstSearch(bst.root)
}

func (bst *BinarySearchTree) breadthFirstSearch(root *Node) []int {
	res := make([]int, 0, bst.Size())

	// Create a queue and insert the root.
	explore := list.New()
	explore.PushFront(root)

	// Loop until the queue is empty.
	for explore.Len() != 0 {
		// Dequeue a node from the queue.
		last := explore.Back()
		node := explore.Back().Value.(*Node)
		explore.Remove(last)

		// Add the child of the extracted node to the queue.
		if node.left != nil {
			explore.PushFront(node.left)
		}
		if node.right != nil {
			explore.PushFront(node.right)
		}

		// Append the extracted node to the result.
		res = append(res, node.data)
	}

	return res
}
