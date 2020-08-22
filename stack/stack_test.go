package stack

import "testing"

func TestPush(t *testing.T) {
	stack := New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	if stack.Size() != 5 {
		t.Errorf("stack should have a size of %d, but it does not", 5)
	}
}

func TestPop(t *testing.T) {
	stack := New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for i := 5; stack.Size() != 0; i-- {
		if v := stack.Pop(); v != i {
			t.Errorf("wrong data: got %d want %d", v, i)
		}
	}
}

func TestPeek(t *testing.T) {
	stack := New()

	assertPanicWithIntReturn(t, stack.Peek)

	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Push(3)
	stack.Pop()

	if v := stack.Peek(); v != 1 {
		t.Errorf("wrong data: got %d want %d", v, 1)
	}
}

func TestContains(t *testing.T) {
	stack := New()

	if stack.Contains(42) {
		t.Errorf("the stack is empty, cannot contains any value")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if !stack.Contains(2) {
		t.Errorf("stack must contains `%d`, but was not found", 2)
	}
	if stack.Contains(5) {
		t.Errorf("stack does not contains `%d`, but was found", 5)
	}
}

func assertPanicWithIntReturn(t *testing.T, f func() int) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic")
		}
	}()
	f()
}
