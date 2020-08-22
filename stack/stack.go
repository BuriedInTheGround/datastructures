package stack

// Stack is a linear data structure that can be accessed only at one end.
type Stack struct {
	data []int
}

// New returns a new Stack instance.
func New() Stack {
	return Stack{data: make([]int, 0)}
}

// Size returns the number of elements inside the Stack.
//
// Complexity: O(1)
func (s *Stack) Size() int {
	return len(s.data)
}

// IsEmpty tells whether the Stack is empty or not.
//
// Complexity: O(1)
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Push adds an element at the top of the Stack.
//
// Complexity: O(1)
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// Pop removes one element from the top of the Stack e returns it.
//
// Complexity: O(1)
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Pop: cannot pop from an empty stack")
	}
	data := s.data[s.Size()-1]
	s.data = s.data[:s.Size()-1]
	return data
}

// Peek returns the value of the element at the top of the Stack.
//
// Complexity: O(1)
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("Peek: cannot peek from an empty stack")
	}
	return s.data[s.Size()-1]
}

// Contains tells whether a specific value is present inside the stack.
//
// Complexity: O(n)
func (s *Stack) Contains(value int) bool {
	if s.IsEmpty() {
		return false
	}
	for _, v := range s.data {
		if v == value {
			return true
		}
	}
	return false
}
