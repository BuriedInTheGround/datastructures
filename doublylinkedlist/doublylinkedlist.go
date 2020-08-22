package doublylinkedlist

import "fmt"

// Node is an element of the DoublyLinkedList, containing a content, a pointer
// to the next node of the chain and one to the previous.
type Node struct {
	content int
	next    *Node
	prev    *Node
}

// Content returns the content value of the node `n`.
func (n *Node) Content() int {
	return n.content
}

// Next returns the next node in the chain with respect to `n`.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns the previous node in the chain with respect to `n`.
func (n *Node) Prev() *Node {
	return n.prev
}

// DoublyLinkedList is a sequential list of Node elements that can be crossed
// in two ways, forward and backward.
type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

// New returns a new DoublyLinkedList instance.
func New() DoublyLinkedList {
	return DoublyLinkedList{head: nil, tail: nil, size: 0}
}

// Head returns the Node which is in front of the list.
func (dll *DoublyLinkedList) Head() *Node {
	return dll.head
}

// Tail returns the Node which is at the end of the list.
func (dll *DoublyLinkedList) Tail() *Node {
	return dll.tail
}

// Size returns the number of elements contained into the list.
func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

// IsEmpty returns whether the list contains at least one element or none.
func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.Size() == 0
}

// Clear removes the list references and makes the size equals to zero.
//
// The GC should take care of removing from memory the elements in between.
func (dll *DoublyLinkedList) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

// InsertFromHead makes a new Node with the given value as content and adds it
// in front of the list.
//
// Complexity: O(1)
func (dll *DoublyLinkedList) InsertFromHead(value int) {
	newNode := Node{content: value, next: nil, prev: nil}
	if dll.IsEmpty() {
		dll.head, dll.tail = &newNode, &newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = &newNode
		dll.head = &newNode
	}
	dll.size++
}

// InsertFromTail makes a new Node with the given value as content and adds it
// at the end of the list.
//
// Complexity: O(1)
func (dll *DoublyLinkedList) InsertFromTail(value int) {
	newNode := Node{content: value, next: nil, prev: nil}
	if dll.IsEmpty() {
		dll.head, dll.tail = &newNode, &newNode
	} else {
		dll.tail.next = &newNode
		newNode.prev = dll.tail
		dll.tail = &newNode
	}
	dll.size++
}

// Contains returns whether the list contains a Node with the given value as
// content or not.
//
// Complexity: O(n)
func (dll *DoublyLinkedList) Contains(value int) bool {
	// If the list is empty there will not be any element.
	if dll.IsEmpty() {
		return false
	}
	for n := dll.Head(); n != nil; n = n.Next() {
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
func (dll *DoublyLinkedList) RemoveFromHead() int {
	if dll.IsEmpty() {
		panic("RemoveFromHead: cannot remove from an empty list")
	}

	data := dll.Head().Content()

	if dll.Head() == dll.Tail() {
		dll.Clear()
		return data
	}

	dll.head = dll.Head().Next()
	dll.head.prev = nil
	dll.size--
	return data
}

// RemoveFromTail removes the tail Node and returns the value that there was
// inside it.
//
// Complexity: O(1)
func (dll *DoublyLinkedList) RemoveFromTail() int {
	if dll.IsEmpty() {
		panic("RemoveFromTail: cannot remove from an empty list")
	}

	data := dll.Tail().Content()

	if dll.Head() == dll.Tail() {
		dll.Clear()
		return data
	}

	dll.tail = dll.Tail().Prev()
	dll.tail.next = nil
	dll.size--
	return data
}

// RemoveFirstOccurrence removes the first occurence of a Node that have a
// content equals to `value`, if exists.
//
// Complexity: O(n)
func (dll *DoublyLinkedList) RemoveFirstOccurrence(value int) error {
	// If the element does not exists inside the list return an error.
	if !dll.Contains(value) {
		return fmt.Errorf("cannot remove value %d not in the list", value)
	}

	// If there is only one Node left, clear the list.
	if dll.Head() == dll.Tail() {
		dll.Clear()
		return nil
	}

	// Otherwise, loop through the chain until the Node next to the cursor is
	// the one that must be removed, and than remove it.
	var cur *Node
	for cur = dll.Head(); cur.Content() != value; cur = cur.Next() {
	}
	cur.prev.next = cur.Next()
	cur.next.prev = cur.Prev()
	dll.size--
	return nil
}

func (dll DoublyLinkedList) String() string {
	res := "[ "
	for n := dll.Head(); n != nil; n = n.Next() {
		res += fmt.Sprintf("%d", n.Content())
		if n.Next() != nil {
			res += ", "
		}
	}
	return res + " ]"
}
