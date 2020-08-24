package queue

import (
	"fmt"

	"github.com/BuriedInTheGround/datastructures/doublylinkedlist"
)

// Queue is a linear data structure that emulate a real queue.
type Queue struct {
	data doublylinkedlist.DoublyLinkedList
}

// New returns a new Queue instance.
func New() Queue {
	return Queue{data: doublylinkedlist.New()}
}

// Size returns the number of elements that are into the queue.
//
// Complexity: O(1)
func (q *Queue) Size() int {
	return q.data.Size()
}

// IsEmpty returns whether the queue is empty or not.
//
// Complexity: O(1)
func (q *Queue) IsEmpty() bool {
	return q.data.IsEmpty()
}

// Enqueue adds an element with the specified `value` to the queue.
//
// Complexity: O(1)
func (q *Queue) Enqueue(value int) {
	q.data.InsertFromHead(value)
}

// Dequeue removes an element from the queue, following the FIFO precedence,
// and returns its value.
//
// Complexity: O(1)
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("Dequeue: cannot dequeue from an empty queue")
	}
	return q.data.RemoveFromTail()
}

// Peek returns the value of the next element that would be dequeued.
//
// Complexity: O(1)
func (q *Queue) Peek() int {
	if q.IsEmpty() {
		panic("Dequeue: cannot peek from an empty queue")
	}
	return q.data.Tail().Content()
}

// Contains returns whether a value is present inside the queue or not.
//
// Complexity: O(n)
func (q *Queue) Contains(value int) bool {
	return q.data.Contains(value)
}

// RemoveFirstOccurrence removes the last element with the specified `value`
// that entered the queue, if exists.
//
// Complexity: O(n)
func (q *Queue) RemoveFirstOccurrence(value int) error {
	err := q.data.RemoveFirstOccurrence(value)
	if err != nil {
		return fmt.Errorf("cannot remove value %d not in queue", value)
	}
	return nil
}
