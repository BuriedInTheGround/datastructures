package linkedlist

import "fmt"

// Node is an element of the LinkedList, containing a content and a pointer to
// the next node of the chain.
type Node struct {
	content int
	next    *Node
}

// Content returns the content value of the node `n`.
func (n *Node) Content() int {
	return n.content
}

// Next returns the next node in the chain with respect to `n`.
func (n *Node) Next() *Node {
	return n.next
}

// LinkedList is a sequential list of Node elements.
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// New return a new LinkedList instance.
func New() LinkedList {
	return LinkedList{head: nil, tail: nil, size: 0}
}

// Head returns the Node which is in front of the list.
func (ll *LinkedList) Head() *Node {
	return ll.head
}

// Tail returns  the Node which is at the end of the list.
func (ll *LinkedList) Tail() *Node {
	return ll.tail
}

// Size returns the number of elements contained into the list.
func (ll *LinkedList) Size() int {
	return ll.size
}

// IsEmpty returns whether the list contains at least one element or none.
func (ll *LinkedList) IsEmpty() bool {
	return ll.Size() == 0
}

// Clear removes the list references and makes the size equals to zero.
//
// The GC should take care of removing from memory the elements in between.
func (ll *LinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

// InsertFromHead makes a new Node with the given value as content and adds it
// in front of the list.
//
// Complexity: O(1)
func (ll *LinkedList) InsertFromHead(value int) {
	newNode := Node{content: value, next: nil}
	if ll.IsEmpty() {
		ll.head, ll.tail = &newNode, &newNode
	} else {
		newNode.next = ll.head
		ll.head = &newNode
	}
	ll.size++
}

// InsertFromTail makes a new Node with the given value as content and adds it
// at the end of the list.
//
// Complexity: O(1)
func (ll *LinkedList) InsertFromTail(value int) {
	newNode := Node{content: value, next: nil}
	if ll.IsEmpty() {
		ll.head, ll.tail = &newNode, &newNode
	} else {
		ll.tail.next = &newNode
		ll.tail = &newNode
	}
	ll.size++
}

// Contains returns whether the list contains a Node with the given value as
// content or not.
//
// Complexity: O(n)
func (ll *LinkedList) Contains(value int) bool {
	// If the list is empty there will not be any element.
	if ll.IsEmpty() {
		return false
	}
	for n := ll.Head(); n != nil; n = n.Next() {
		if n.Content() == value {
			return true
		}
	}
	return false
}

// RemoveFromHead removes the head Node and returns the value that there was
// inside it.
//
// Complexity: O(1)
func (ll *LinkedList) RemoveFromHead() int {
	if ll.IsEmpty() {
		panic("RemoveFromHead: cannot remove from an empty list")
	}
	data := ll.head.content
	ll.head = ll.head.next
	ll.size--
	return data
}

// RemoveFromTail removes the tail Node and returns the value that there was
// inside it.
//
// Complexity: O(n)
func (ll *LinkedList) RemoveFromTail() int {
	if ll.IsEmpty() {
		panic("RemoveFromTail: cannot remove from an empty list")
	}

	data := ll.tail.content

	// If there is only one Node left, clear the list and return data.
	if ll.Head() == ll.Tail() {
		ll.Clear()
		return data
	}

	// Otherwise, loop through the chain until the pre-tail Node, change the
	// tail, and return data.
	var cur *Node
	for cur = ll.Head(); cur.Next() != ll.Tail(); cur = cur.Next() {
	}
	ll.tail = cur
	ll.tail.next = nil
	ll.size--

	return data
}

// RemoveFirstOccurrence removes a Node that have a content equals to `value`,
// if exists.
//
// Complexity: O(n)
func (ll *LinkedList) RemoveFirstOccurrence(value int) error {
	// If the element does not exists inside the list return an error.
	if !ll.Contains(value) {
		return fmt.Errorf("cannot remove value %d not in the list", value)
	}

	// If there is only one Node left, clear the list.
	if ll.Head() == ll.Tail() {
		ll.Clear()
		return nil
	}

	// Otherwise, loop through the chain until the Node next to the cursor is
	// the one that must be removed, and than remove it.
	var cur *Node
	for cur = ll.Head(); cur.Next().Content() != value; cur = cur.Next() {
	}
	cur.next = cur.Next().Next()
	return nil
}

func (ll LinkedList) String() string {
	res := "[ "
	for n := ll.Head(); n != nil; n = n.Next() {
		res += fmt.Sprintf("%d", n.Content())
		if n.Next() != nil {
			res += ", "
		}
	}
	return res + " ]"
}
