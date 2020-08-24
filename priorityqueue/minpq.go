package priorityqueue

import "fmt"

// MinPQ is an abstract data type (ADT) that works in the same way as a Queue
// but every element has a priority that determines the order in which it is
// extracted.
//
// This implementation removes the element with the least priority first.
type MinPQ struct {
	data []int
}

// New returns a new MinPQ instance.
func New() MinPQ {
	return MinPQ{data: make([]int, 0)}
}

// Size returns the number of elements that are into the PQ.
//
// Complexity: O(1)
func (pq *MinPQ) Size() int {
	return len(pq.data)
}

// IsEmpty returns whether the PQ is empty or not.
//
// Complexity: O(1)
func (pq *MinPQ) IsEmpty() bool {
	return pq.Size() == 0
}

// Add adds a new element with the specified `value` to the PQ.
//
// Complexity: O(log(n))
func (pq *MinPQ) Add(value int) {
	pq.data = append(pq.data, value)
	pq.bubbleUp(pq.Size() - 1)
}

// RemoveMin removes an element from the PQ, following the priority order, and
// returns its value.
//
// Complexity: O(log(n))
func (pq *MinPQ) RemoveMin() int {
	if pq.IsEmpty() {
		panic("RemoveMin: cannot remove from an empty PQ")
	}
	res := pq.data[0]
	pq.data[0] = pq.data[pq.Size()-1]
	pq.data = pq.data[:pq.Size()-1]
	pq.bubbleDown(0)
	return res
}

// Peek returns the value of the next element the would be returned by
// RemoveMin.
//
// Complexity: O(1)
func (pq *MinPQ) Peek() int {
	if pq.IsEmpty() {
		panic("Peek: cannot peek from an empty PQ")
	}
	return pq.data[0]
}

// Contains returns whether the PQ contains the specified `value` or not.
//
// Complexity: O(n)
func (pq *MinPQ) Contains(value int) bool {
	if pq.IsEmpty() {
		return false
	}
	for _, v := range pq.data {
		if v == value {
			return true
		}
	}
	return false
}

// RemoveFirstOccurrence removes the first occurrence of the specified `value`.
//
// Complexity: O(n)
func (pq *MinPQ) RemoveFirstOccurrence(value int) error {
	if !pq.Contains(value) {
		return fmt.Errorf("cannot remove value %d not in PQ", value)
	}

	// Find the index of the element to remove.
	var i int
	for i = range pq.data {
		if pq.data[i] == value {
			break
		}
	}

	// Swap the element with the last.
	last := pq.data[pq.Size()-1]
	pq.data[i] = last
	pq.data = pq.data[:pq.Size()-1]

	// If the removed element was the last element there is no need to bubble.
	if i == pq.Size() {
		return nil
	}

	// Bubble down first, then, if nothing happened, try to bubble up.
	pq.bubbleDown(i)
	if pq.data[i] == last {
		pq.bubbleUp(i)
	}
	return nil
}

func (pq *MinPQ) bubbleUp(fromIndex int) {
	// Loop until a parent exists and it has a greater priority.
	for fromIndex > 0 && pq.data[fromIndex] < pq.data[(fromIndex-1)/2] {
		pq.swap(fromIndex, (fromIndex-1)/2)
		fromIndex = (fromIndex - 1) / 2
	}
}

func (pq *MinPQ) bubbleDown(fromIndex int) {
	// Loop until a left child exists.
	for (2*fromIndex)+1 < pq.Size() {
		// Find which child, left or right, has the least priority.
		minChildIndex := (2 * fromIndex) + 1
		if (2*fromIndex)+2 < pq.Size() && pq.data[(2*fromIndex)+2] < pq.data[minChildIndex] {
			minChildIndex = (2 * fromIndex) + 2
		}

		// Bubble down if the child has a smaller priority.
		if pq.data[minChildIndex] < pq.data[fromIndex] {
			pq.swap(minChildIndex, fromIndex)
		} else {
			// Otherwise we are done.
			return
		}

		fromIndex = minChildIndex
	}
}

func (pq *MinPQ) swap(idx1, idx2 int) {
	temp := pq.data[idx1]
	pq.data[idx1] = pq.data[idx2]
	pq.data[idx2] = temp
}

func (pq MinPQ) String() string {
	res := "[ "
	for _, v := range pq.data {
		res += fmt.Sprintf("%d", v)
		res += " "
	}
	return res + "]"
}
